package models

import (
	"BudgBackend/src/database"
	"fmt"
)

//type Report interface {
//	GetMonthlyReport(userId int, months int) (Report, error)
//	GetCategoryReport(userId int, days int) (Report, error)
//}

type Report struct {
	Labels []string  `json:"labels"`
	Data   []float64 `json:"data"`
}

func (m *Report) GetMonthlyReport(userId int, months int) (Report, error) {
	InfoLogger.Println("Getting monthly report")
	var report Report
	query := fmt.Sprintf("SELECT MONTH(t.date), SUM(t.amount) FROM User_transaction t INNER JOIN Budget b ON t.budget_id = b.id WHERE b.user_id = %d and t.type='expense' GROUP BY MONTH(t.date) order by MONTH(t.date) LIMIT %d", userId, months)

	rows, err := database.QueryDB(query)
	if err != nil {
		ErrorLogger.Println("Error getting monthly report: ", err)
		return report, err
	}
	i := 0
	for rows.Next() {
		i++
		var month int
		var amount float64
		err = rows.Scan(&month, &amount)
		if err != nil {
			ErrorLogger.Println("Error scanning monthly report: ", err)
		}
		report.Labels = append(report.Labels, getMonthName(month))

		report.Data = append(report.Data, amount)
	}

	if i == 0 {
		return report, fmt.Errorf("No transactions for user with id: %d", userId)
	}
	return report, nil

}

func (m *Report) GetCategoryReport(userId int, days int) (Report, error) {
	InfoLogger.Println("Getting category report")
	var report Report
	query := fmt.Sprintf("SELECT c.name, SUM(e.amount) FROM User_transaction e INNER JOIN Category c ON e.category_id = c.id WHERE e.date BETWEEN DATE_SUB(NOW(), INTERVAL %d DAY) AND NOW() AND e.type='expense' AND e.user_id=%d GROUP BY c.name order by SUM(e.amount)", days, userId)

	rows, err := database.QueryDB(query)
	if err != nil {
		ErrorLogger.Println("Error getting category report: ", err)
		return report, err
	}
	i := 0
	for rows.Next() {
		i++
		var category string
		var amount float64
		err = rows.Scan(&category, &amount)
		if err != nil {
			ErrorLogger.Println("Error scanning category report: ", err)
		}
		report.Labels = append(report.Labels, category)
		report.Data = append(report.Data, amount)
	}

	if i == 0 {
		return report, fmt.Errorf("No transactions for user with id: %d", userId)
	}
	return report, nil

}

func getMonthName(monthNumber int) string {
	switch monthNumber {
	case 1:
		return "Jan"
	case 2:
		return "Feb"
	case 3:
		return "Mar"
	case 4:
		return "Apr"
	case 5:
		return "May"
	case 6:
		return "Jun"
	case 7:
		return "Jul"
	case 8:
		return "Aug"
	case 9:
		return "Sep"
	case 10:
		return "Oct"
	case 11:
		return "Nov"
	case 12:
		return "Dec"
	}
	return ""
}
