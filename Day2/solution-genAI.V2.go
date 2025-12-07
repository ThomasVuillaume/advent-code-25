package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

// --- Mêmes fonctions logiques qu'avant ---

func isPeriodic(s string) bool {
	n := len(s)
	if n < 2 {
		return false
	}
	for length := 1; length <= n/2; length++ {
		if n%length != 0 {
			continue
		}
		pattern := s[:length]
		repeats := n / length
		if s == strings.Repeat(pattern, repeats) {
			return true
		}
	}
	return false
}

func processRange(rangeStr string) int {
	parts := strings.Split(rangeStr, "-")
	if len(parts) != 2 {
		return 0 //format de plage invalide: %s", rangeStr
	}

	min, _ := strconv.Atoi(parts[0])
	max, _ := strconv.Atoi(parts[1])

	sum := 0
	for i := min; i <= max; i++ {
		if isPeriodic(strconv.Itoa(i)) {
			sum += i
		}
	}
	return sum
}

// --- La partie Concurrence ---

func main() {
	// 1. Configuration
	// On récupère le nombre de cœurs CPU disponibles (ex: 8 sur un Mac M1, 4 sur un i5, etc.)
	numWorkers := runtime.NumCPU()
	fmt.Printf("Démarrage avec %d workers en parallèle...\n", numWorkers)

	// Canaux de communication
	// 'jobs' transporte les chaînes "100-200"
	jobs := make(chan string, 100)
	// 'results' transporte les sommes partielles calculées
	results := make(chan int, 100)

	// WaitGroup sert à attendre que tous les ouvriers aient fini
	var wg sync.WaitGroup

	// 2. Lancement des Workers (Ouvriers)
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// L'ouvrier écoute le canal 'jobs' tant qu'il est ouvert
			for rangeString := range jobs {
				// Il fait le travail lourd
				partialSum := processRange(rangeString)
				// Il envoie le résultat
				results <- partialSum
			}
		}(w)
	}

	// 3. Le Distributeur (Lecture du fichier)
	// On le lance dans une goroutine pour ne pas bloquer la lecture des résultats
	go func() {
		file, err := os.Open("input.txt")
		if err != nil {
			log.Fatalf("Erreur fichier: %v", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			ranges := strings.Split(line, ",")
			for _, r := range ranges {
				// On envoie le travail dans le canal
				jobs <- r
			}
		}
		// Très important : on ferme le canal 'jobs' pour dire aux ouvriers
		// "Il n'y a plus de travail, rentrez chez vous".
		close(jobs)
	}()

	// 4. Le Gestionnaire de Fermeture
	// Une goroutine qui attend que tous les workers aient fini (wg.Wait)
	// pour fermer le canal de résultats.
	go func() {
		wg.Wait()
		close(results)
	}()

	// 5. La Récolte (Main thread)
	// Le thread principal additionne les résultats au fur et à mesure qu'ils arrivent
	totalSol := 0
	for pSum := range results {
		totalSol += pSum
	}

	fmt.Println("Solution Finale :", totalSol)
}
