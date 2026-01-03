package utils

const specialityDbName = "speciality_db"

var GoogleSheetsConfig = GoogleSheets{
	Levels:             []string{"Рівень вищої освіти", "levels"},
	Statuses:           []string{"Статус дисципліни", "statuses"},
	Departments:        []string{"Кафедри", "departments"},
	Specialties:        []string{"Спеціальності", "specialties"},
	Programs:           []string{"Освітні програми", "programs"},
	DepartmentPrograms: []string{"Освітні програми кафедр", "department_programs"},
	Forms:              []string{"Форма навчання", "forms"},
}
