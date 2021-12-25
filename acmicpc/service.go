package acmicpc

import (
	"acmicpc_checker_v2_backend/db"
	"acmicpc_checker_v2_backend/model"
	"acmicpc_checker_v2_backend/solvedInfo"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func getWebInfo(url string) *goquery.Document {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func GetStatus(targetURL string) (map[int]time.Time, string) {
	data := map[int]time.Time{}

	doc := getWebInfo(targetURL)
	doc.Find("#status-table tbody tr").Each(func(i int, s *goquery.Selection) {
		problem_id_string := s.Find(".problem_title").First().Text()
		problem_id, _ := strconv.Atoi(problem_id_string)

		time_string, _ := s.Find(".show-date").First().Attr("data-timestamp")
		time_int, _ := strconv.Atoi(time_string)
		findTimestamp := time.Unix(int64(time_int), 0)

		data[problem_id] = findTimestamp
	})

	nextURL, exists := doc.Find("#next_page").Attr("href")

	if exists {
		return data, nextURL
	}

	return data, ""
}

func GetProblemListFromExerciseBook(id int) string {
	doc := getWebInfo(fmt.Sprintf("https://www.acmicpc.net/workbook/view/%d", id))

	problemList := make([]string, 0)

	doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		val, _ := s.Find("td").First().Html()
		problemList = append(problemList, val)
	})

	return strings.Join(problemList, ",")
}

func GetTotalSolvedCount(userID string) int {
	doc := getWebInfo(fmt.Sprintf("https://www.acmicpc.net/user/%s", userID))
	solvedCountString := doc.Find("#u-solved").Text()
	solvedCount, err := strconv.Atoi(solvedCountString)

	if err != nil {
		log.Printf("%s의 문제 해결 갯수를 가져오는데 실패하였습니다.\n", userID)
		return -1
	}

	return solvedCount
}

func SolvedCheker(userID string) {
	db := db.Dbconnect()
	url := fmt.Sprintf("https://www.acmicpc.net/status?user_id=%s&result_id=4", userID)

	for {
		log.Printf("%s\n", url)
		data, nextURL := GetStatus(url)

		// 추가된 것이 한개라도 있다면, 0이 아닌 값이 나온다.
		insertCount := 0
		for k := range data {
			s := &model.SolvedInfo{ProblemID: k, SolvedTime: data[k], StudentAID: userID}

			if solvedInfo.Create(db, s) == nil {
				insertCount++
			}
		}

		if insertCount > 0 && nextURL != "" {
			url = "https://www.acmicpc.net" + nextURL
		} else {
			break
		}
	}
}

func CheckedAcmipcData(userIDList []string, checkList []int) {
	result := FindCheckedData(userIDList, checkList)

	conv := map[string][]int{}

	for _, r := range result {
		conv[r.StudentAID] = append(conv[r.StudentAID], r.ProblemID)
	}

	for _, userID := range userIDList {
		if len(checkList) != len(conv[userID]) {
			SolvedCheker(userID)
		}
	}
}

func FindCheckedData(acmicpc_id_list []string, problem_list []int) []model.SolvedInfo {
	db := db.Dbconnect()

	var result []model.SolvedInfo
	db.Model(&model.SolvedInfo{}).Order("student_a_id, problem_id").Where("student_a_id in ? and problem_id in ?", acmicpc_id_list, problem_list).Find(&result)

	return result
}

func SolvedData(assignmentID uint) []model.SolvedInfo {
	db := db.Dbconnect()
	type tmp struct {
		ClassInfoID   uint
		ProblemList   string
		AcmicpcIDList string
	}

	var t tmp

	db.Raw("SELECT assigninfo.class_info_id, problem_list, group_concat(acmicpc_id) as acmicpc_id_list FROM (SELECT class_info_id, problem_list FROM assignments WHERE id = ?) as assigninfo JOIN class_students ON assigninfo.class_info_id = class_students.class_info_id JOIN students s ON s.id = class_students.student_id GROUP BY assigninfo.class_info_id", assignmentID).Scan(&t)

	acmicpc_ids := t.AcmicpcIDList
	problems := t.ProblemList

	acmicpc_id_list := make([]string, 0)
	problem_list := make([]int, 0)

	acmicpc_id_list = append(acmicpc_id_list, strings.Split(acmicpc_ids, ",")...)

	for _, problem := range strings.Split(problems, ",") {
		problem_int, _ := strconv.Atoi(problem)
		problem_list = append(problem_list, problem_int)
	}
	CheckedAcmipcData(acmicpc_id_list, problem_list)

	result := FindCheckedData(acmicpc_id_list, problem_list)
	return result
}
