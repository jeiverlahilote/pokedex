package main

import "fmt"

const NMAX = 1000
const nBox = 32
const box = 30

type Pokemon struct {
	Number           int
	Name             string
	Level            int
	Gender           string
	Type             string
	TeraType         string
	Stats            stat
	Ability          string
	HeldItem         string
	Moves            string
	Lokasi           string
	waktuPenangkapan int
}

type stat struct {
	HP        int
	Attack    int
	Defense   int
	Speed     int
	SpAttack  int
	SpDefense int
}

var PCbox [nBox][box]Pokemon

func intro() {
	fmt.Println("Welcome to POKEDEX")
	fmt.Scanln()

}

func menu(pilihan int, pokemon Pokemon, noBox, noSlot int) {
	for {
		fmt.Println("=========================")
		fmt.Println("        MENU        ")
		fmt.Println("=========================")
		fmt.Println(" 1.     INVENTORY     ")
		fmt.Println(" 2.    MENAMBAHKAN    ")
		fmt.Println(" 3.     MENGHAPUS     ")
		fmt.Println(" 4.   Cari Pokemon    ")
		fmt.Println(" 5.      Keluar       ")
		fmt.Println("=========================")
		fmt.Println("Pilihan anda")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			displayInventory(pilihan, pokemon, noBox, noSlot)
		} else if pilihan == 2 {
			menambahPokemon(pokemon, noBox, noSlot)
		} else if pilihan == 3 {
			menghapusPokemon(noBox, noSlot)
		} else if pilihan == 4 {
			menuSearch()
		} else if pilihan == 5{
		    return
		} else {
		    fmt.Println("Pilihan tidak valid")
		}
	}
}

func displayInventory(pilihan int, pokemon Pokemon, noBox, noSlot int) {
	var i, j, n int
	fmt.Println("=========================")
	fmt.Println("    INVENTORY    ")
	fmt.Println("=========================")
	fmt.Println("Sort by:")
	fmt.Println(" 1. Tanggal Penangkapan")
	fmt.Println(" 2. Tera Type")
	fmt.Println(" 3. Type Pokemon")
	fmt.Println(" 4. Keluar")
	fmt.Println("=========================")
	fmt.Println("Pilihan anda")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		sortByTanggalPenangkapan(&i, &j, &n)
	} else if pilihan == 2 {
		sortByTeraType(&i, &j, &n)
	} else if pilihan == 3 {
		sortByTypePokemon(&i, &j, &n)
	} else if pilihan == 4 {
		menu(pilihan, pokemon, noBox, noSlot)
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func sortByTanggalPenangkapan(i, j, n *int) {
    /*insertionSort*/
	fmt.Println("=========================")
	fmt.Println(" SORT BY TANGGAL PENANGKAPAN ")
	fmt.Println("=========================")

	var pokemonList [NMAX]Pokemon
	*n = flattenPCBox(&pokemonList)

	for *i = 1; *i < *n; *i++ {
		key := pokemonList[*i]
		*j = *i - 1

		for *j >= 0 && pokemonList[*j].waktuPenangkapan > key.waktuPenangkapan {
			pokemonList[*j+1] = pokemonList[*j]
			*j--
		}

		pokemonList[*j+1] = key
	}

	displayPokemonList(&pokemonList)
}

func sortByTeraType(i, j, n *int) {
    /*selectionSort*/
	fmt.Println("=========================")
	fmt.Println("     SORT BY TERA TYPE     ")
	fmt.Println("=========================")

	var pokemonList [NMAX]Pokemon
	*n = flattenPCBox(&pokemonList)

	for *i = 0; *i < *n-1; *i++ {
		minIndex := *i

		for *j = *i + 1; *j < *n; *j++ {
			if pokemonList[*j].TeraType < pokemonList[minIndex].TeraType {
				minIndex = *j
			}
		}

		if minIndex != *i {
			pokemonList[*i], pokemonList[minIndex] = pokemonList[minIndex], pokemonList[*i]
		}
	}

	displayPokemonList(&pokemonList)
}

func sortByTypePokemon(i, j, n *int) {
    /*selectionSort*/
	fmt.Println("=========================")
	fmt.Println("     SORT BY TYPE     ")
	fmt.Println("=========================")

	var pokemonList [NMAX]Pokemon
	*n = flattenPCBox(&pokemonList)

	for *i = 0; *i < *n-1; *i++ {
		minIndex := *i

		for *j = *i + 1; *j < *n; *j++ {
			if pokemonList[*j].Type < pokemonList[minIndex].Type {
				minIndex = *j
			}
		}

		if minIndex != *i {
			pokemonList[*i], pokemonList[minIndex] = pokemonList[minIndex], pokemonList[*i]
		}
	}

	displayPokemonList(&pokemonList)
}

func menambahPokemon(pokemon Pokemon, noBox, noSlot int) {
	fmt.Println("=========================")
	fmt.Println("     MENAMBAHKAN POKEMON     ")
	fmt.Println("=========================")

	fmt.Println("Pilih Box (1-32)")
	fmt.Scan(&noBox)
	fmt.Println("Pilih slot (1-30)")
	fmt.Scan(&noSlot)

	if noBox < 1 || noBox > nBox || noSlot < 1 || noSlot > box {
		fmt.Println("Box atau slot tidak valid")
		return
	}

	if PCbox[noBox-1][noSlot-1].Name != "" {
		fmt.Println("Slot sudah terisi")
		return
	}

	fmt.Println("Nomor Pokemon:")
	fmt.Scan(&pokemon.Number)
	fmt.Println("Nama Pokemon:")
	fmt.Scan(&pokemon.Name)
	fmt.Println("Level Pokemon:")
	fmt.Scan(&pokemon.Level)
	fmt.Println("Gender Pokemon:")
	fmt.Scan(&pokemon.Gender)
	fmt.Println("Type Pokemon:")
	fmt.Scan(&pokemon.Type)
	fmt.Println("Tera Type Pokemon:")
	fmt.Scan(&pokemon.TeraType)
	fmt.Println("Stats Pokemon:")
	fmt.Println("HP:")
	fmt.Scan(&pokemon.Stats.HP)
	fmt.Println("Attack:")
	fmt.Scan(&pokemon.Stats.Attack)
	fmt.Println("Defense:")
	fmt.Scan(&pokemon.Stats.Defense)
	fmt.Println("Speed:")
	fmt.Scan(&pokemon.Stats.Speed)
	fmt.Println("Sp. Attack:")
	fmt.Scan(&pokemon.Stats.SpAttack)
	fmt.Println("Sp. Defense:")
	fmt.Scan(&pokemon.Stats.SpDefense)
	fmt.Println("Ability Pokemon:")
	fmt.Scan(&pokemon.Ability)
	fmt.Println("Held Item Pokemon:")
	fmt.Scan(&pokemon.HeldItem)
	fmt.Println("Moves Pokemon:")
	fmt.Scan(&pokemon.Moves)
	fmt.Println("Lokasi Penangkapan Pokemon:")
	fmt.Scan(&pokemon.Lokasi)
	fmt.Println("Tahun Penangkapan:")
	fmt.Scan(&pokemon.waktuPenangkapan)

	PCbox[noBox-1][noSlot-1] = pokemon
	fmt.Println("Pokemon berhasil ditambahkan")
}

func menghapusPokemon(noBox, noSlot int) {
	fmt.Println("=========================")
	fmt.Println("     MENGHAPUS POKEMON     ")
	fmt.Println("=========================")

	fmt.Println("Pilih Box (1-32)")
	fmt.Scan(&noBox)
	fmt.Println("Pilih slot (1-30)")
	fmt.Scan(&noSlot)

	if noBox < 1 || noBox > nBox || noSlot < 1 || noSlot > box {
		fmt.Println("Box atau slot tidak valid")
		return
	}

	if PCbox[noBox-1][noSlot-1].Name == "" {
		fmt.Println("Slot kosong")
		return
	}

	PCbox[noBox-1][noSlot-1] = Pokemon{}
	fmt.Println("Pokemon berhasil dihapus")
}

func flattenPCBox(pokemonList *[NMAX]Pokemon) int {
    /* menggabungkan semua Pokemon dari PCbox (kotak penyimpanan)
    menjadi satu daftar yang terurut*/
	index := 0
	for i := 0; i < nBox; i++ {
		for j := 0; j < box; j++ {
			if PCbox[i][j].Name != "" {
				pokemonList[index] = PCbox[i][j]
				index++
			}
		}
	}
	return index
}

func menuSearch() {
	var pilihan int
	fmt.Println("=========================")
	fmt.Println("    PENCARIAN    ")
	fmt.Println("=========================")
	fmt.Println("Pencarian Dari:")
	fmt.Println(" 1. Nomor Pokemon")
	fmt.Println(" 2. Level Pokemon")
	fmt.Println(" 3. Nama Pokemon")
	fmt.Println("=========================")
	fmt.Println("Pilihan anda")
	fmt.Scan(&pilihan)

	var no, level int
	var nama string
	if pilihan == 1 {
		fmt.Println("Nomor Pokemon:")
		fmt.Scan(&no)
		foundPokemon := cariPokemonByNumber(no)
		if foundPokemon != nil {
			displayPokemonList(&[NMAX]Pokemon{*foundPokemon})
		} else {
			fmt.Println("Pokemon tidak ditemukan")
		}
	} else if pilihan == 2 {
		fmt.Println("Level Pokemon:")
		fmt.Scan(&level)
		foundPokemon := cariPokemonByLevel(level)
		if foundPokemon != nil {
			displayPokemonList(&[NMAX]Pokemon{*foundPokemon})
		} else {
			fmt.Println("Pokemon tidak ditemukan")
		}
	} else if pilihan == 3 {
	    fmt.Println("Nama Pokemon")
	    fmt.Scan(&nama)
	    foundPokemon := cariPokemonByNama(nama)
		if foundPokemon != nil {
			displayPokemonList(&[NMAX]Pokemon{*foundPokemon})
		} else {
		   fmt.Println("Pokemon tidak ditemukan") 
		}
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func cariPokemonByNumber(number int) *Pokemon {
    /*SeqSearch*/
    var i int 
    var j int
    i = 0
	for i < nBox {
	    j = 0
		for j < box {
			if PCbox[i][j].Number == number {
				return &PCbox[i][j]
			}
			j++
		}
		i++
	}
	return nil
}

func cariPokemonByLevel(Level int) *Pokemon {
    /*SeqSearch*/
    var j int
    var i int
    i = 0
    for i < nBox {
        j = 0
        for j < box {
            if PCbox[i][j].Level == Level {
                return &PCbox[i][j]
            }
            j++
        }
        i++
    }
    return nil
}

func cariPokemonByNama(nama string) *Pokemon {
    /*SeqSearch*/
    var j int
    var i int
    i = 0
    for i < nBox {
        j = 0
        for j < box {
            if PCbox[i][j].Name == nama {
                return &PCbox[i][j]
            }
            j++
        }
        i++
    }
    return nil
}


func displayPokemonList(pokemonList *[NMAX]Pokemon) {
	fmt.Println("=========================")
	fmt.Println("     DAFTAR POKEMON     ")
	fmt.Println("=========================")
	for i := 0; i < len(pokemonList); i++ {
		pokemon := pokemonList[i]
		if pokemon.Name != "" {
			fmt.Println("Nomor Pokemon:", pokemon.Number)
			fmt.Println("Nama Pokemon:", pokemon.Name)
			fmt.Println("Level Pokemon:", pokemon.Level)
			fmt.Println("Gender Pokemon:", pokemon.Gender)
			fmt.Println("Type Pokemon:", pokemon.Type)
			fmt.Println("Tera Type Pokemon:", pokemon.TeraType)
			fmt.Println("Stats Pokemon:")
			fmt.Println("HP:", pokemon.Stats.HP)
			fmt.Println("Attack:", pokemon.Stats.Attack)
			fmt.Println("Defense:", pokemon.Stats.Defense)
			fmt.Println("Speed:", pokemon.Stats.Speed)
			fmt.Println("Sp. Attack:", pokemon.Stats.SpAttack)
			fmt.Println("Sp. Defense:", pokemon.Stats.SpDefense)
			fmt.Println("Ability Pokemon:", pokemon.Ability)
			fmt.Println("Held Item Pokemon:", pokemon.HeldItem)
			
			fmt.Println("Moves Pokemon:", pokemon.Moves)
			fmt.Println("Lokasi Penangkapan Pokemon:", pokemon.Lokasi)
			fmt.Println("Waktu Penangkapan Pokemon:", pokemon.waktuPenangkapan)
			fmt.Println("=========================")
		}
	}
}

func main() {
	var pilihan, noBox, noSlot int
	var pokemon Pokemon
	intro()
	menu(pilihan, pokemon, noBox, noSlot)
}