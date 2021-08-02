package main

import (
	"math"
	"sort"
)

//amount of recommendations that will be returned to the frontend service
const amountOfRecommendations = 6

//allocate memory for future calculation of the simmilarity score
var similarityScore = make(map[SubjectsCsvStruct]float64)

type Recommender struct{}

// calculatioon of cosine similarity based on two input arrays
// first is the annotated subejct which will be comapred against the query
// second one represents the client send user preference
func (reco *Recommender) CosineSimilarity(a []float64, b []float64) (cosine float64, err error) {
	count := 0
	lengthA := len(a)
	lengthB := len(b)
	if lengthA > lengthB {
		count = lengthA
	} else {
		count = lengthB
	}
	sumA := 0.0
	s1 := 0.0
	s2 := 0.0
	for k := 0; k < count; k++ {
		if k >= lengthA {
			s2 += math.Pow(b[k], 2)
			continue
		}
		if k >= lengthB {
			s1 += math.Pow(a[k], 2)
			continue
		}
		sumA += a[k] * b[k]
		s1 += math.Pow(a[k], 2)
		s2 += math.Pow(b[k], 2)
	}
	return sumA / (math.Sqrt(s1) * math.Sqrt(s2)), nil
}

//convert a Userpreference object to list of floats usable for CosineSimilarity similarity calculation
func (reco *Recommender) convertToFloat(preferences UserPreferences) []float64 {
	return []float64{float64(preferences.SOFTWARE), float64(preferences.AI), float64(preferences.LOWLEVEL),
		float64(preferences.SECURITY), float64(preferences.WEB), float64(preferences.THEORETICAL)}
}

//convert query and data to specific format for cosin similarity and calculate this
func (reco *Recommender) calculateSimilarity(subject SubjectsCsvStruct, query UserPreferences) {
	subjectConverted := reco.convertToFloat(subject.Preferences)
	//primes := convertToFloat(preferences)
	queryConverted := reco.convertToFloat(query)
	score, _ := reco.CosineSimilarity(subjectConverted, queryConverted)

	similarityScore[subject] = score

}

// provide recommendations based on the preference user vector
// get subjects from mariaDb and compares the vecotrs based on a similarity score
func (reco *Recommender) provideRecommendation(preferences UserPreferences) []SubjectsCsvStruct {
	//using CSV file
	//csvData := readSubjectsCSV()
	//allRecords := fillStructWithCsvData(csvData)

	//using Maria DB to get all subjects that are available
	allRecords := getAllSubjectsFromMariaDb()

	for _, subject := range allRecords {
		reco.calculateSimilarity(subject, preferences)
	}

	sortedSimilarityScore := sortMapByValue(similarityScore)
	recommendations := sortedSimilarityScore[:amountOfRecommendations]
	recommendationsWithoutSimScore := recommendations.GetKeys(recommendations)

	return recommendationsWithoutSimScore

}

//sorting map by value and returns a sorted by value map
func sortMapByValue(wordFrequencies map[SubjectsCsvStruct]float64) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

//implementing structure for traversing through pairs subjects
type Pair struct {
	Key   SubjectsCsvStruct
	Value float64
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) GetKeys(list PairList) []SubjectsCsvStruct {
	var keysSliice []SubjectsCsvStruct
	for _, pairListElem := range list {
		keysSliice = append(keysSliice, pairListElem.Key)
	}
	return keysSliice
}
