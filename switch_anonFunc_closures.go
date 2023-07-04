package main

func main() {

}

func prediction(dayOfWeek string) string {
	switch dayOfWeek {
	case "пн":
		return ""
	case "вт":
		return ""
	case "ср":
		return ""
	case "чт":
		return "н"
	default:
		return "невалидный день недели"
	}
}
