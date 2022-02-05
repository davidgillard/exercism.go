// fichier : lasagna.go
// auteur : David GILLARD
// date : 05/02/2022

package lasagna
// TODO: define the 'OvenTime' constant
const OvenTime = 40
// TODO: define the 'RemainingOvenTime()' function
func RemainingOvenTime(usedTime int) int {
    return OvenTime - usedTime
}
// TODO: define the 'PreparationTime()' function
func PreparationTime(layers int) int {
    return layers * 2
}
// TODO: define the 'ElapsedTime()' function
func ElapsedTime(layers int, time int) int {
    return PreparationTime(layers) + time
}
