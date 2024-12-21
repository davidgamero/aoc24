package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const EMPTYID = -1

func ExpandDiskMap(rawDiskMap string) ([]int, error) {
	var expandedDiskMap []int
	isData := true
	for iRawDiskMap, r := range rawDiskMap {
		fileID := iRawDiskMap / 2
		numberOfBlocks, err := strconv.Atoi(string(r))
		if err != nil {
			return expandedDiskMap, fmt.Errorf("converting raw disk map rune: %w", err)
		}

		for iBlock := 0; iBlock < numberOfBlocks; iBlock += 1 {
			if isData {
				expandedDiskMap = append(expandedDiskMap, fileID)
			} else {
				expandedDiskMap = append(expandedDiskMap, EMPTYID)
			}
		}
		isData = !isData
	}
	return expandedDiskMap, nil
}

func CompactExpandedMap(expandedMap []int) []int {
	iFirstEmpty := 0
	iLastData := len(expandedMap) - 1
	for expandedMap[iLastData] == '.' {
		iLastData -= 1
	}
	compactedMap := []int{}
	for iFirstEmpty < iLastData {
		// copy blocks with existing data
		if expandedMap[iFirstEmpty] != EMPTYID {
			compactedMap = append(compactedMap, expandedMap[iFirstEmpty])
			iFirstEmpty += 1
			continue
		}
		if expandedMap[iLastData] == EMPTYID {
			iLastData -= 1
			continue
		}
		compactedMap = append(compactedMap, expandedMap[iLastData])
		iFirstEmpty += 1
		iLastData -= 1
	}

	for len(compactedMap) < len(expandedMap) {
		compactedMap = append(compactedMap, EMPTYID)
	}

	return compactedMap
}

func GetFirstEmptySpace(expandedMap []int, spaceLength int) (int, bool) {
	var emptyLen int
	for i, fileID := range expandedMap {
		if fileID != EMPTYID {
			emptyLen = 0
			continue
		}
		emptyLen += 1
		if emptyLen == spaceLength {
			return i - spaceLength + 1, true
		}
	}
	return -1, false
}

// Swap n entries between [i:i+length] and [j:j+length]
func SwapLen(s []int, i, j, n int) ([]int, error) {
	swapped := make([]int, len(s))
	lenCopied := copy(swapped, s)
	if lenCopied != len(s) {
		return swapped, fmt.Errorf("failed to copy")
	}
	toi := s[j : j+n]
	if len(toi) != n {
		return swapped, fmt.Errorf("failed to extract proper toi block length")
	}
	toj := s[i : i+n]
	if len(toj) != n {
		return swapped, fmt.Errorf("failed to extract proper toj block length")
	}
	swapped = slices.Replace(swapped, i, i+n, toi...)
	swapped = slices.Replace(swapped, j, j+n, toj...)
	return swapped, nil
}

func CompactWholeFiles(expandedMap []int) []int {
	blockLen := 0
	compactedMap := make([]int, len(expandedMap))
	copy(compactedMap, expandedMap)
	//movedIDs := map[int]bool{}
	for iBlockRead := len(expandedMap) - 1; iBlockRead > 0; iBlockRead -= 1 {
		blockLen += 1
		fmt.Println("iBlockRead=", iBlockRead, " blockLen=", blockLen)
		isBlockStart := expandedMap[iBlockRead-1] != expandedMap[iBlockRead] && expandedMap[iBlockRead] != EMPTYID
		if expandedMap[iBlockRead-1] != expandedMap[iBlockRead] {
			blockLen = 1
		}

		if !isBlockStart {
			fmt.Println(" not block start, continuing [iBlockRead-1:]=", expandedMap[iBlockRead-1:])
			continue
		}
		fmt.Println(" block start at i=", iBlockRead, " for fileID=", expandedMap[iBlockRead], " with blocklen=", blockLen)
		// is block start, attempt placement
		iPlacementTarget, ok := GetFirstEmptySpace(compactedMap, blockLen)
		if ok && iPlacementTarget < iBlockRead {
			// swap with target
			fmt.Println("  swapping found empty space at i=", iPlacementTarget, " for blocklen=", blockLen)
			swapped, err := SwapLen(compactedMap, iPlacementTarget, iBlockRead, blockLen)
			if err != nil {
				panic(err)
			}
			compactedMap = swapped
		}
		fmt.Println(compactedMap)
		blockLen = 0
	}
	return compactedMap
}

func GetChecksum(expandedMap []int) (int, error) {
	checksum := 0
	for i, fileID := range expandedMap {
		if fileID != EMPTYID {
			checksum += i * fileID
		}
	}
	return checksum, nil
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	rawMap := string(lines[0])

	expandedMap, err := ExpandDiskMap(rawMap)
	if err != nil {
		panic(err)
	}
	fmt.Println("expanded map:")
	fmt.Println(expandedMap)

	compactedMap := CompactExpandedMap(expandedMap)
	fmt.Println("compacted map:")
	fmt.Println(compactedMap)

	checksum, err := GetChecksum(compactedMap)

	fmt.Printf("p1 checksum: %d\n", checksum)
}
