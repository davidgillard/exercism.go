/* file: lasagna_master.go
  author: David Gillard
  date: 24/02/2022
  http://patorjk.com/software/taag


 o                                                                                     o          o                            o
<|>                                                                                   <|\        /|>                          <|>
/ \                                                                                   / \\o    o// \                          < >
\o/           o__ __o/      __o__   o__ __o/    o__ __o/  \o__ __o      o__ __o/      \o/ v\  /v \o/     o__ __o/      __o__   |        o__  __o   \o__ __o
 |           /v     |      />  \   /v     |    /v     |    |     |>    /v     |        |   <\/>   |     /v     |      />  \    o__/_   /v      |>   |     |>
/ \         />     / \     \o     />     / \  />     / \  / \   / \   />     / \      / \        / \   />     / \     \o       |      />      //   / \   < >
\o/         \      \o/      v\    \      \o/  \      \o/  \o/   \o/   \      \o/      \o/        \o/   \      \o/      v\      |      \o    o/     \o/
 |           o      |        <\    o      |    o      |    |     |     o      |        |          |     o      |        <\     o       v\  /v __o   |
/ \ _\o__/_  <\__  / \  _\o__</    <\__  / \   <\__  < >  / \   / \    <\__  / \      / \        / \    <\__  / \  _\o__</     <\__     <\/> __/>  / \
                                                      |
                                              o__     o
                                              <\__ __/>
*/

package lasagna

// TODO: define the 'PreparationTime()' function

func PreparationTime(layers []string, ptpl int) int {
	switch {
	case ptpl == 0:
		return len(layers) * 2
	default:
		return len(layers) * ptpl
	}
}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, l := range layers {
		switch {
		case l == "noodles":
			noodles += 50
		case l == "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	//return append(myList, friendsList[len(friendsList)-1])
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe(listIngredient []float64, portions int) []float64 {
	newList := make([]float64, len(listIngredient))
	for i, ingredient := range listIngredient {
		newList[i] = ingredient / 2 * float64(portions)
	}
	return newList
}
