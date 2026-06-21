package response

import "fmt"

func PrintLine() {
	fmt.Println("--------------------------------------------------")
}

func PrintConfig(input string, recursive bool, workers int) {
	PrintLine()
	fmt.Println("CONFIG")
	PrintLine()
	fmt.Printf("Input Folder : %s\n", input)
	fmt.Printf("Recursive    : %t\n", recursive)
	fmt.Printf("Workers      : %d\n", workers)
	fmt.Println()
}

func PrintSummary(folderCount, fileCount, totalImages int) {
	PrintLine()
	fmt.Println("SUMMARY")
	PrintLine()
	fmt.Printf("Folders      : %d\n", folderCount)
	fmt.Printf("Files        : %d\n", fileCount)
	fmt.Printf("Total Images : %d\n", totalImages)
}

func PrintImageCounts(counts map[string]int) int {
	total := 0

	PrintLine()
	fmt.Println("IMAGE COUNTS")
	PrintLine()

	for ext, count := range counts {
		fmt.Printf("%-6s : %d\n", ext, count)
		total += count
	}

	return total
}