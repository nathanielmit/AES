package main

import "fmt"


var Sbox = [16][16]int{
    { 0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76 } ,
    { 0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0, 0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0 } ,
    { 0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc, 0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15 } ,
    { 0x04, 0xc7, 0x23, 0xc3, 0x18, 0x96, 0x05, 0x9a, 0x07, 0x12, 0x80, 0xe2, 0xeb, 0x27, 0xb2, 0x75 } ,
    { 0x09, 0x83, 0x2c, 0x1a, 0x1b, 0x6e, 0x5a, 0xa0, 0x52, 0x3b, 0xd6, 0xb3, 0x29, 0xe3, 0x2f, 0x84 } ,
    { 0x53, 0xd1, 0x00, 0xed, 0x20, 0xfc, 0xb1, 0x5b, 0x6a, 0xcb, 0xbe, 0x39, 0x4a, 0x4c, 0x58, 0xcf } ,
    { 0xd0, 0xef, 0xaa, 0xfb, 0x43, 0x4d, 0x33, 0x85, 0x45, 0xf9, 0x02, 0x7f, 0x50, 0x3c, 0x9f, 0xa8 } ,
    { 0x51, 0xa3, 0x40, 0x8f, 0x92, 0x9d, 0x38, 0xf5, 0xbc, 0xb6, 0xda, 0x21, 0x10, 0xff, 0xf3, 0xd2 } ,
    { 0xcd, 0x0c, 0x13, 0xec, 0x5f, 0x97, 0x44, 0x17, 0xc4, 0xa7, 0x7e, 0x3d, 0x64, 0x5d, 0x19, 0x73 } ,
    { 0x60, 0x81, 0x4f, 0xdc, 0x22, 0x2a, 0x90, 0x88, 0x46, 0xee, 0xb8, 0x14, 0xde, 0x5e, 0x0b, 0xdb } ,
    { 0xe0, 0x32, 0x3a, 0x0a, 0x49, 0x06, 0x24, 0x5c, 0xc2, 0xd3, 0xac, 0x62, 0x91, 0x95, 0xe4, 0x79 } ,
    { 0xe7, 0xc8, 0x37, 0x6d, 0x8d, 0xd5, 0x4e, 0xa9, 0x6c, 0x56, 0xf4, 0xea, 0x65, 0x7a, 0xae, 0x08 } ,
    { 0xba, 0x78, 0x25, 0x2e, 0x1c, 0xa6, 0xb4, 0xc6, 0xe8, 0xdd, 0x74, 0x1f, 0x4b, 0xbd, 0x8b, 0x8a } ,
    { 0x70, 0x3e, 0xb5, 0x66, 0x48, 0x03, 0xf6, 0x0e, 0x61, 0x35, 0x57, 0xb9, 0x86, 0xc1, 0x1d, 0x9e } ,
    { 0xe1, 0xf8, 0x98, 0x11, 0x69, 0xd9, 0x8e, 0x94, 0x9b, 0x1e, 0x87, 0xe9, 0xce, 0x55, 0x28, 0xdf } ,
    { 0x8c, 0xa1, 0x89, 0x0d, 0xbf, 0xe6, 0x42, 0x68, 0x41, 0x99, 0x2d, 0x0f, 0xb0, 0x54, 0xbb, 0x16 }}

var Rcon = [...]int{ 0x00000000,
    0x01000000, 0x02000000, 0x04000000, 0x08000000,
    0x10000000, 0x20000000, 0x40000000, 0x80000000,
    0x1B000000, 0x36000000, 0x6C000000, 0xD8000000,
    0xAB000000, 0x4D000000, 0x9A000000, 0x2F000000,
    0x5E000000, 0xBC000000, 0x63000000, 0xC6000000,
    0x97000000, 0x35000000, 0x6A000000, 0xD4000000,
    0xB3000000, 0x7D000000, 0xFA000000, 0xEF000000,
    0xC5000000, 0x91000000, 0x39000000, 0x72000000,
    0xE4000000, 0xD3000000, 0xBD000000, 0x61000000,
    0xC2000000, 0x9F000000, 0x25000000, 0x4A000000,
    0x94000000, 0x33000000, 0x66000000, 0xCC000000,
    0x83000000, 0x1D000000, 0x3A000000, 0x74000000,
    0xE8000000, 0xCB000000, 0x8D000000}

var InvSbox = [16][16]int {
    { 0x52, 0x09, 0x6a, 0xd5, 0x30, 0x36, 0xa5, 0x38, 0xbf, 0x40, 0xa3, 0x9e, 0x81, 0xf3, 0xd7, 0xfb } ,
    { 0x7c, 0xe3, 0x39, 0x82, 0x9b, 0x2f, 0xff, 0x87, 0x34, 0x8e, 0x43, 0x44, 0xc4, 0xde, 0xe9, 0xcb } ,
    { 0x54, 0x7b, 0x94, 0x32, 0xa6, 0xc2, 0x23, 0x3d, 0xee, 0x4c, 0x95, 0x0b, 0x42, 0xfa, 0xc3, 0x4e } ,
    { 0x08, 0x2e, 0xa1, 0x66, 0x28, 0xd9, 0x24, 0xb2, 0x76, 0x5b, 0xa2, 0x49, 0x6d, 0x8b, 0xd1, 0x25 } ,
    { 0x72, 0xf8, 0xf6, 0x64, 0x86, 0x68, 0x98, 0x16, 0xd4, 0xa4, 0x5c, 0xcc, 0x5d, 0x65, 0xb6, 0x92 } ,
    { 0x6c, 0x70, 0x48, 0x50, 0xfd, 0xed, 0xb9, 0xda, 0x5e, 0x15, 0x46, 0x57, 0xa7, 0x8d, 0x9d, 0x84 } ,
    { 0x90, 0xd8, 0xab, 0x00, 0x8c, 0xbc, 0xd3, 0x0a, 0xf7, 0xe4, 0x58, 0x05, 0xb8, 0xb3, 0x45, 0x06 } ,
    { 0xd0, 0x2c, 0x1e, 0x8f, 0xca, 0x3f, 0x0f, 0x02, 0xc1, 0xaf, 0xbd, 0x03, 0x01, 0x13, 0x8a, 0x6b } ,
    { 0x3a, 0x91, 0x11, 0x41, 0x4f, 0x67, 0xdc, 0xea, 0x97, 0xf2, 0xcf, 0xce, 0xf0, 0xb4, 0xe6, 0x73 } ,
    { 0x96, 0xac, 0x74, 0x22, 0xe7, 0xad, 0x35, 0x85, 0xe2, 0xf9, 0x37, 0xe8, 0x1c, 0x75, 0xdf, 0x6e } ,
    { 0x47, 0xf1, 0x1a, 0x71, 0x1d, 0x29, 0xc5, 0x89, 0x6f, 0xb7, 0x62, 0x0e, 0xaa, 0x18, 0xbe, 0x1b } ,
    { 0xfc, 0x56, 0x3e, 0x4b, 0xc6, 0xd2, 0x79, 0x20, 0x9a, 0xdb, 0xc0, 0xfe, 0x78, 0xcd, 0x5a, 0xf4 } ,
    { 0x1f, 0xdd, 0xa8, 0x33, 0x88, 0x07, 0xc7, 0x31, 0xb1, 0x12, 0x10, 0x59, 0x27, 0x80, 0xec, 0x5f } ,
    { 0x60, 0x51, 0x7f, 0xa9, 0x19, 0xb5, 0x4a, 0x0d, 0x2d, 0xe5, 0x7a, 0x9f, 0x93, 0xc9, 0x9c, 0xef } ,
    { 0xa0, 0xe0, 0x3b, 0x4d, 0xae, 0x2a, 0xf5, 0xb0, 0xc8, 0xeb, 0xbb, 0x3c, 0x83, 0x53, 0x99, 0x61 } ,
    { 0x17, 0x2b, 0x04, 0x7e, 0xba, 0x77, 0xd6, 0x26, 0xe1, 0x69, 0x14, 0x63, 0x55, 0x21, 0x0c, 0x7d }}

func printState(state [4][4]int) {
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
           fmt.Printf("%02x", state[j][i])
            if (j+1)%16==0 {
                fmt.Println()
            }
        }
    }
    fmt.Printf("\n")
}

func getKey(expandedKey []int, keyIndex int, Nk int) [4][4]int {
    //cols := (Nk+7)*4
    var key [4][4]int
    for i:=0; i<4; i++ { //cols
        for j:=0; j<4; j++ { //rows
            key[j][i] = expandedKey[(keyIndex*16)+(4*i)+j]
        }
    }
    return key
}

func ffadd(a int, b int) int {
    return a^b
}

func xtime(val int) int {
    var mask int = 0x80
    if (mask & val) > 0 {
        return ((val<<1) ^ 0x1b) & 0xff
    }
    return ((val << 1) & 0xff)
}

func ffMultiply(a int, b int) int {
    var mask int = 0x1
    var vals [10]int
    var addVals []int

    vals[0] = a
    var sum int = 0

    for i := 1; i < 8; i++ {
        vals[i] = xtime(vals[i-1])
    }
    for i := 0; i < 8; i++ {
        if (mask & b) != 0 {
            addVals = append(addVals, vals[i])
        }
        mask = mask << 1
    }
    for i := 0; i < len(addVals); i++ {
        sum ^= addVals[i]
    }
    return sum
}

func subWord(word [4]int) [4]int {
    for i := 0; i < 4; i++ {
        word[i] = Sbox[(word[i] >> 4) & 0x0f][word[i] & 0x0f]
    }
    return word
}

func rotWord(word [4]int) [4]int {
    temp := word[0]
    word[0] = word[1]
    word[1] = word[2]
    word[2] = word[3]
    word[3] = temp;
    return word
}

func printWord(word [4]int) {
    for p:=0; p<4; p++ {
        fmt.Printf("%x ", word[p])
    }
    fmt.Println()
}

func keyExpansion(key [][]int, Nk int) []int {
    cols := (Nk+7)*4
    var expandedKey = make([]int, cols*4)
    var rconIndex int = 1
    var temp [4]int
    // copy initial keys
    for columns := 0; columns < Nk; columns++ { //rows
        for rows := 0; rows < 4; rows++ { //columns
            expandedKey[(4*columns) + rows] = key[rows][columns]
        }
    }

    Nb := 4
    Nr := Nk+6
    for i:=Nk; i < Nb*(Nr+1); i++ { //columns
        for z:=0; z<4; z++ { //get the previous column
            temp[z] = expandedKey[(4*(i-1))+z]
        }
        if i % Nk == 0 {
            temp = subWord(rotWord(temp))
            temp[0] = temp[0] ^ (Rcon[rconIndex]>>24)
            rconIndex++
        } else if Nk > 6 && i%Nk==4 {
            temp = subWord(temp)
        }
        for z:=0; z<4; z++ {
            expandedKey[(i*4)+z] = expandedKey[((i-Nk)*4)+z] ^ temp[z]
        }
    }
    return expandedKey

}

func addRoundKey(state [4][4]int, roundKey [4][4]int) [4][4]int {
    var roundedState [4][4]int
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            roundedState[j][i] = state[j][i] ^ roundKey[j][i]
        }
    }
    return roundedState
}

//Substitute bytes of input array with bytes from Sbox
func subBytes(input [4][4]int) [4][4]int {
    var subbedArray [4][4]int
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            s := input[i][j] & 0xf
            f := input[i][j] >> 4
            //fmt.Printf("%v, %v, %x, %x\n", f, s, input[i][j],Sbox[f][s])
            subbedArray[i][j] = Sbox[f][s]
        }
    }
    //fmt.Println("\n\n")
    return subbedArray
}

func invSubBytes(input [4][4]int) [4][4]int {
    var subbedArray [4][4]int
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            s := input[i][j] & 0xf
            f := input[i][j] >> 4
            subbedArray[i][j] = InvSbox[f][s]
        }
    }
    //fmt.Println("\n\n")
    return subbedArray
}

func invShiftRows(input [4][4]int) [4][4]int {
    shiftedArray := input

    shiftedArray[1][0] = input[1][3]
    shiftedArray[1][1] = input[1][0]
    shiftedArray[1][2] = input[1][1]
    shiftedArray[1][3] = input[1][2]

    shiftedArray[2][0] = input[2][2]
    shiftedArray[2][1] = input[2][3]
    shiftedArray[2][2] = input[2][0]
    shiftedArray[2][3] = input[2][1]

    shiftedArray[3][0] = input[3][1]
    shiftedArray[3][1] = input[3][2]
    shiftedArray[3][2] = input[3][3]
    shiftedArray[3][3] = input[3][0]

    return shiftedArray
}

func invMixColumns(state [4][4]int) [4][4]int {
    var newState [4][4]int
    for i := 0; i < 4; i++ {
        newState[0][i] = ffMultiply(0x0e, state[0][i]) ^ ffMultiply(0x0b, state[1][i]) ^ ffMultiply(0x0d, state[2][i]) ^ ffMultiply(0x09, state[3][i])
        newState[1][i] = ffMultiply(0x09, state[0][i]) ^ ffMultiply(0x0e, state[1][i]) ^ ffMultiply(0x0b, state[2][i]) ^ ffMultiply(0x0d, state[3][i])
        newState[2][i] = ffMultiply(0x0d, state[0][i]) ^ ffMultiply(0x09, state[1][i]) ^ ffMultiply(0x0e, state[2][i]) ^ ffMultiply(0x0b,state[3][i])
        newState[3][i] = ffMultiply(0x0b, state[0][i]) ^ ffMultiply(0x0d, state[1][i]) ^ ffMultiply(0x09, state[2][i]) ^ ffMultiply(0x0e,state[3][i])
    }
    return newState

}

func shiftRows(input [4][4]int) [4][4]int {
    shiftedArray := input

    shiftedArray[1][0] = input[1][1]
    shiftedArray[1][1] = input[1][2]
    shiftedArray[1][2] = input[1][3]
    shiftedArray[1][3] = input[1][0]

    shiftedArray[2][0] = input[2][2]
    shiftedArray[2][1] = input[2][3]
    shiftedArray[2][2] = input[2][0]
    shiftedArray[2][3] = input[2][1]

    shiftedArray[3][0] = input[3][3]
    shiftedArray[3][1] = input[3][0]
    shiftedArray[3][2] = input[3][1]
    shiftedArray[3][3] = input[3][2]

    return shiftedArray
}

func mixColumns(state [4][4]int) [4][4]int {
    var newState [4][4]int
    for i := 0; i < 4; i++ {
        newState[0][i] = ffMultiply(0x02, state[0][i]) ^ ffMultiply(0x03, state[1][i]) ^ state[2][i] ^ state[3][i]
        newState[1][i] = state[0][i] ^ ffMultiply(0x02, state[1][i]) ^ ffMultiply(0x03, state[2][i]) ^ state[3][i]
        newState[2][i] = state[0][i] ^ state[1][i] ^ ffMultiply(0x02, state[2][i]) ^ ffMultiply(0x03,state[3][i])
        newState[3][i] = ffMultiply(0x03, state[0][i]) ^ state[1][i] ^ state[2][i] ^ ffMultiply(0x02,state[3][i])
    }
    return newState
}

func encrypt(state [4][4]int, key [][]int, Nk int) [4][4]int {
    var roundKey [4][4]int
    var expandedKey []int
    expandedKey = keyExpansion(key, Nk)
    roundKey = getKey(expandedKey, 0, Nk)
    fmt.Printf("round[0].input:   ")
    printState(state)
    fmt.Printf("round[0].k_sch:   ")
    printState(roundKey)
    state = addRoundKey(state, roundKey)
    for i := 1; i < Nk+6; i++ { //loop this Nk+6 times
      fmt.Printf("round[%v].start:   ", i)
      printState(state)

      state = subBytes(state)
      fmt.Printf("round[%v].s_box:   ", i)
      printState(state)

      state = shiftRows(state)
      fmt.Printf("round[%v].s_row:   ", i)
      printState(state)

      state = mixColumns(state)
      fmt.Printf("round[%v].m_col:   ", i)
      printState(state)

      fmt.Printf("round[%v].k_sch:   ", i)
      printState(getKey(expandedKey, i, Nk))
      state = addRoundKey(state, getKey(expandedKey, i, Nk))
    }
    fmt.Printf("round[%v].start:   ", Nk+6)
    printState(state)

    state = subBytes(state)
    fmt.Printf("round[%v].s_box:   ", Nk+6)
    printState(state)

    state = shiftRows(state)
    fmt.Printf("round[%v].s_row:   ", Nk+6)
    printState(state)

    fmt.Printf("round[%v].k_sch:   ", Nk+6)
    printState(getKey(expandedKey, Nk+6, Nk))
    state = addRoundKey(state, getKey(expandedKey, Nk+6, Nk))

    fmt.Printf("round[%v].output:  ", Nk+6)
    printState(state)
    return state
}

func decrypt(state [4][4]int, key [][]int, Nk int) [4][4]int {
    Nr := Nk+6
    var roundKey [4][4]int
    var expandedKey []int
    expandedKey = keyExpansion(key, Nk)
    roundKey = getKey(expandedKey, Nr, Nk)

    fmt.Printf("round[0].iinput:   ")
    printState(state)

    fmt.Printf("round[0].ik_sch:   ")
    printState(roundKey)

    state = addRoundKey(state, roundKey)
    for round:=(Nr-1); round>0; round-- {
        fmt.Printf("round[%v].istart:   ", Nr-round)
        printState(state)

        state = invShiftRows(state)
        fmt.Printf("round[%v].is_row:   ", Nr-round)
        printState(state)

        state =invSubBytes(state)
        fmt.Printf("round[%v].is_box:   ", Nr-round)
        printState(state)

        roundKey = getKey(expandedKey, round, Nk)
        fmt.Printf("round[%v].ik_sch:   ", Nr-round)
        printState(roundKey)

        state = addRoundKey(state, roundKey)
        fmt.Printf("round[%v].ik_add:   ", Nr-round)
        printState(state)

        state = invMixColumns(state)
    }

    fmt.Printf("round[%v].istart:   ", Nr)
    printState(state)

    state = invShiftRows(state)
    fmt.Printf("round[%v].is_row:   ", Nr)
    printState(state)

    state = invSubBytes(state)
    fmt.Printf("round[%v].is_box:   ", Nr)
    printState(state)

    roundKey = getKey(expandedKey, 0, Nk)
    fmt.Printf("round[%v].ik_sch:   ", Nr)
    printState(roundKey)

    state = addRoundKey(state, roundKey)
    fmt.Printf("round[%v].ioutput:  ", Nr)
    printState(state)

    return state
}

func main() {
/*
    testArray := [4][4]int{{0x32, 0x88, 0x31, 0xe0},
                           {0x43, 0x5a, 0x31, 0x37},
                           {0xf6, 0x30, 0x98, 0x07},
                           {0xa8, 0x8d, 0xa2, 0x34}}
*/

    input := [4][4]int{{0x00, 0x44, 0x88, 0xcc},
                           {0x11, 0x55, 0x99, 0xdd},
                           {0x22, 0x66, 0xaa, 0xee},
                           {0x33, 0x77, 0xbb, 0xff}}

    key_128 := [][]int{{0x00, 0x04, 0x08, 0x0c},
                     {0x01, 0x05, 0x09, 0x0d},
                     {0x02, 0x06, 0x0a, 0x0e},
                     {0x03, 0x07, 0x0b, 0x0f}}

    key_192 := [][]int{{0x00, 0x04, 0x08, 0x0c, 0x10, 0x14},
                       {0x01, 0x05, 0x09, 0x0d, 0x11, 0x15},
                       {0x02, 0x06, 0x0a, 0x0e, 0x12, 0x16},
                       {0x03, 0x07, 0x0b, 0x0f, 0x13, 0x17}}

    key_256 := [][]int{{0x00, 0x04, 0x08, 0x0c, 0x10, 0x14, 0x18, 0x1c},
                       {0x01, 0x05, 0x09, 0x0d, 0x11, 0x15, 0x19, 0x1d},
                       {0x02, 0x06, 0x0a, 0x0e, 0x12, 0x16, 0x1a, 0x1e},
                       {0x03, 0x07, 0x0b, 0x0f, 0x13, 0x17, 0x1b, 0x1f}}

    //TEST AES-128
    fmt.Println("AES-128")
    Nk := 4
    fmt.Printf("PLAINTEXT: ")
    printState(input)
    fmt.Printf("KEY: ")
    for column := 0; column < 4; column++ {
        for row := 0; row < 4; row++ {
            fmt.Printf("%02x", key_128[row][column])
        }
    }
    fmt.Println()
    fmt.Println("CIPHER (ENCRYPT):")
    cypher := encrypt(input, key_128, Nk)
    fmt.Println()
    fmt.Println("INVERSE CIPHER (DECRYPT):")
    decrypt(cypher, key_256, Nk)
    fmt.Printf("\n\n")

    //TEST AES-192
    fmt.Println("AES-192")
    Nk = 6
    fmt.Printf("PLAINTEXT: ")
    printState(input)
    fmt.Printf("KEY: ")
    for column := 0; column < 6; column++ {
        for row := 0; row < 4; row++ {
            fmt.Printf("%02x", key_192[row][column])
        }
    }
    fmt.Println()
    fmt.Println("CIPHER (ENCRYPT):")
    cypher = encrypt(input, key_192, Nk)
    fmt.Println()
    fmt.Println("INVERSE CIPHER (DECRYPT):")
    decrypt(cypher, key_256, Nk)
    fmt.Printf("\n")

    //TEST AES-256
    fmt.Println()
    fmt.Println("AES-256")
    Nk = 8
    fmt.Printf("PLAINTEXT: ")
    printState(input)
    fmt.Printf("KEY: ")
    for column := 0; column < 8; column++ {
        for row := 0; row < 4; row++ {
            fmt.Printf("%02x", key_256[row][column])
        }
    }
    fmt.Println()
    fmt.Println("CIPHER (ENCRYPT):")
    cypher = encrypt(input, key_256, Nk)
    fmt.Println()
    fmt.Println("INVERSE CIPHER (DECRYPT):")
    decrypt(cypher, key_256, Nk)
}
