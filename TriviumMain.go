package main

import (
	"time"
	"math/rand"
	"fmt"
)
var key[80]byte
var IV[80]byte
var LFSR[288]byte
func genKey() {
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<80;i++ {
		key[i] = randomInt(0, 2);
		IV[i] = randomInt(0, 2);
	}
	for  i:=0; i<288;i++ {
		LFSR[i] = 0
	}
}
func randomInt(min, max int) byte {
	return byte(min + rand.Intn(max-min))
}
func keySetup() {
	var m int = 0
	for i := 0; i < 93; i++ {
		if i > 79 {
			LFSR[0] = 0
		} else {
			LFSR[i] = key[i]
		}
	}
	for i := 93; i < 177; i++ {
		if i > 172 {
			LFSR[i] = 0
		} else {
			LFSR[i] = IV[m]
		}
		m++
	}
	for i := 178; i < 288; i++ {
		if i > 285 {
			LFSR[i] = 1
		} else {
			LFSR[1] = 0
		}
	}
	for i := 0; i < (4 * 288); i++ {
		var t1 byte = LFSR[65] ^ LFSR[90]&LFSR[91] ^ LFSR[92] ^ LFSR[93]
		var t2 byte = LFSR[161] ^ LFSR[174]&LFSR[175] ^ LFSR[176] ^ LFSR[163]
		var t3 byte = LFSR[242] ^ LFSR[285]&LFSR[186] ^ LFSR[287] ^ LFSR[68]
		for i := 91; i >= 0; i-- {
			if i == 0 {
				LFSR[1] = LFSR[0]
				LFSR[0] = t3
			} else {
				LFSR[i+1] = LFSR[i]
			}
		}
		for i := 175; i >= 93; i-- {
			if i == 93 {
				LFSR[94] = LFSR[93]
				LFSR[93] = t1
			} else {
				LFSR[i+1] = LFSR[i]
			}
		}
		for i := 286; i >= 177; i-- {
			if i == 177 {
				LFSR[178] = LFSR[177];
				LFSR[177] = t2;
			} else {
				LFSR[i+1] = LFSR[i];
			}
		}
	}
}
func genKeyStream(){
	var t1 = (byte) (LFSR[65] ^ LFSR[92])
	var t2 = (byte) (LFSR[161] ^ LFSR[176])
	var t3 = (byte) (LFSR[242] ^ LFSR[287])
	var z = (byte) (t1 ^ t2 ^ t3)
	t1 = (byte) (t1 ^ LFSR[90] & LFSR[91] ^ LFSR[170])
	t2 = (byte) (t2 ^ LFSR[174] & LFSR[175] ^ LFSR[263])
	t3 = (byte) (t3 ^ LFSR[285] & LFSR[286] ^ LFSR[68])

	LFSR[92] = LFSR[91]
	LFSR[91] = LFSR[90]
	LFSR[90] = LFSR[89]
	LFSR[89] = LFSR[88]
	LFSR[88] = LFSR[87]
	LFSR[87] = LFSR[86]
	LFSR[86] = LFSR[85]
	LFSR[85] = LFSR[84]
	LFSR[84] = LFSR[83]
	LFSR[83] = LFSR[82]
	LFSR[82] = LFSR[81]
	LFSR[81] = LFSR[80]
	LFSR[80] = LFSR[79]
	LFSR[79] = LFSR[78]
	LFSR[78] = LFSR[77]
	LFSR[77] = LFSR[76]
	LFSR[76] = LFSR[75]
	LFSR[75] = LFSR[74]
	LFSR[74] = LFSR[73]
	LFSR[73] = LFSR[72]
	LFSR[72] = LFSR[71]
	LFSR[71] = LFSR[70]
	LFSR[70] = LFSR[69]
	LFSR[69] = LFSR[68]
	LFSR[68] = LFSR[67]
	LFSR[67] = LFSR[66]
	LFSR[66] = LFSR[65]
	LFSR[65] = LFSR[64]
	LFSR[64] = LFSR[63]
	LFSR[63] = LFSR[62]
	LFSR[62] = LFSR[61]
	LFSR[61] = LFSR[60]
	LFSR[60] = LFSR[59]
	LFSR[59] = LFSR[58]
	LFSR[58] = LFSR[57]
	LFSR[57] = LFSR[56]
	LFSR[56] = LFSR[55]
	LFSR[55] = LFSR[54]
	LFSR[54] = LFSR[53]
	LFSR[53] = LFSR[52]
	LFSR[52] = LFSR[51]
	LFSR[51] = LFSR[50]
	LFSR[50] = LFSR[49]
	LFSR[49] = LFSR[48]
	LFSR[48] = LFSR[47]
	LFSR[47] = LFSR[46]
	LFSR[46] = LFSR[45]
	LFSR[45] = LFSR[44]
	LFSR[44] = LFSR[43]
	LFSR[43] = LFSR[42]
	LFSR[42] = LFSR[41]
	LFSR[41] = LFSR[40]
	LFSR[40] = LFSR[39]
	LFSR[39] = LFSR[38]
	LFSR[38] = LFSR[37]
	LFSR[37] = LFSR[36]
	LFSR[36] = LFSR[35]
	LFSR[35] = LFSR[34]
	LFSR[34] = LFSR[33]
	LFSR[33] = LFSR[32]
	LFSR[32] = LFSR[31]
	LFSR[31] = LFSR[30]
	LFSR[30] = LFSR[29]
	LFSR[29] = LFSR[28]
	LFSR[28] = LFSR[27]
	LFSR[27] = LFSR[26]
	LFSR[26] = LFSR[25]
	LFSR[25] = LFSR[24]
	LFSR[24] = LFSR[23]
	LFSR[23] = LFSR[22]
	LFSR[22] = LFSR[21]
	LFSR[21] = LFSR[20]
	LFSR[20] = LFSR[19]
	LFSR[19] = LFSR[18]
	LFSR[18] = LFSR[17]
	LFSR[17] = LFSR[16]
	LFSR[16] = LFSR[15]
	LFSR[15] = LFSR[14]
	LFSR[14] = LFSR[13]
	LFSR[13] = LFSR[12]
	LFSR[12] = LFSR[11]
	LFSR[11] = LFSR[10]
	LFSR[10] = LFSR[9]
	LFSR[9] = LFSR[8]
	LFSR[8] = LFSR[7]
	LFSR[7] = LFSR[6]
	LFSR[6] = LFSR[5]
	LFSR[5] = LFSR[4]
	LFSR[4] = LFSR[3]
	LFSR[3] = LFSR[2]
	LFSR[2] = LFSR[1]
	LFSR[1] = LFSR[0]
	LFSR[0] = t3

	LFSR[176] = LFSR[175]
	LFSR[175] = LFSR[174]
	LFSR[174] = LFSR[173]
	LFSR[173] = LFSR[172]
	LFSR[172] = LFSR[171]
	LFSR[171] = LFSR[170]
	LFSR[170] = LFSR[169]
	LFSR[169] = LFSR[168]
	LFSR[168] = LFSR[167]
	LFSR[167] = LFSR[166]
	LFSR[166] = LFSR[165]
	LFSR[165] = LFSR[164]
	LFSR[164] = LFSR[163]
	LFSR[163] = LFSR[162]
	LFSR[162] = LFSR[161]
	LFSR[161] = LFSR[160]
	LFSR[160] = LFSR[159]
	LFSR[159] = LFSR[158]
	LFSR[158] = LFSR[157]
	LFSR[157] = LFSR[156]
	LFSR[156] = LFSR[155]
	LFSR[155] = LFSR[154]
	LFSR[154] = LFSR[153]
	LFSR[153] = LFSR[152]
	LFSR[152] = LFSR[151]
	LFSR[151] = LFSR[150]
	LFSR[150] = LFSR[149]
	LFSR[149] = LFSR[148]
	LFSR[148] = LFSR[147]
	LFSR[147] = LFSR[146]
	LFSR[146] = LFSR[145]
	LFSR[145] = LFSR[144]
	LFSR[144] = LFSR[143]
	LFSR[143] = LFSR[142]
	LFSR[142] = LFSR[141]
	LFSR[141] = LFSR[140]
	LFSR[140] = LFSR[139]
	LFSR[139] = LFSR[138]
	LFSR[138] = LFSR[137]
	LFSR[136] = LFSR[135]
	LFSR[135] = LFSR[134]
	LFSR[134] = LFSR[133]
	LFSR[133] = LFSR[132]
	LFSR[132] = LFSR[131]
	LFSR[131] = LFSR[130]
	LFSR[130] = LFSR[129]
	LFSR[129] = LFSR[128]
	LFSR[128] = LFSR[127]
	LFSR[127] = LFSR[126]
	LFSR[126] = LFSR[125]
	LFSR[125] = LFSR[124]
	LFSR[124] = LFSR[123]
	LFSR[123] = LFSR[122]
	LFSR[122] = LFSR[121]
	LFSR[121] = LFSR[120]
	LFSR[120] = LFSR[119]
	LFSR[119] = LFSR[118]
	LFSR[118] = LFSR[117]
	LFSR[117] = LFSR[116]
	LFSR[116] = LFSR[115]
	LFSR[115] = LFSR[114]
	LFSR[114] = LFSR[113]
	LFSR[113] = LFSR[112]
	LFSR[112] = LFSR[111]
	LFSR[111] = LFSR[110]
	LFSR[110] = LFSR[109]
	LFSR[109] = LFSR[108]
	LFSR[108] = LFSR[107]
	LFSR[107] = LFSR[106]
	LFSR[106] = LFSR[105]
	LFSR[105] = LFSR[104]
	LFSR[104] = LFSR[103]
	LFSR[103] = LFSR[102]
	LFSR[102] = LFSR[101]
	LFSR[101] = LFSR[100]
	LFSR[100] = LFSR[99]
	LFSR[99] = LFSR[98]
	LFSR[98] = LFSR[97]
	LFSR[97] = LFSR[96]
	LFSR[96] = LFSR[95]
	LFSR[95] = LFSR[94]
	LFSR[94] = LFSR[93]
	LFSR[93] = t1

	LFSR[287] = LFSR[286]
	LFSR[286] = LFSR[285]
	LFSR[285] = LFSR[284]
	LFSR[284] = LFSR[283]
	LFSR[283] = LFSR[282]
	LFSR[282] = LFSR[281]
	LFSR[281] = LFSR[280]
	LFSR[280] = LFSR[279]
	LFSR[279] = LFSR[278]
	LFSR[278] = LFSR[277]
	LFSR[277] = LFSR[276]
	LFSR[276] = LFSR[275]
	LFSR[275] = LFSR[274]
	LFSR[274] = LFSR[273]
	LFSR[273] = LFSR[272]
	LFSR[272] = LFSR[271]
	LFSR[271] = LFSR[270]
	LFSR[270] = LFSR[269]
	LFSR[269] = LFSR[268]
	LFSR[268] = LFSR[267]
	LFSR[267] = LFSR[266]
	LFSR[266] = LFSR[265]
	LFSR[265] = LFSR[264]
	LFSR[264] = LFSR[263]
	LFSR[263] = LFSR[262]
	LFSR[262] = LFSR[261]
	LFSR[261] = LFSR[260]
	LFSR[260] = LFSR[259]
	LFSR[259] = LFSR[258]
	LFSR[258] = LFSR[257]
	LFSR[257] = LFSR[256]
	LFSR[256] = LFSR[255]
	LFSR[255] = LFSR[254]
	LFSR[254] = LFSR[253]
	LFSR[253] = LFSR[252]
	LFSR[252] = LFSR[251]
	LFSR[251] = LFSR[250]
	LFSR[250] = LFSR[249]
	LFSR[249] = LFSR[248]
	LFSR[248] = LFSR[247]
	LFSR[247] = LFSR[246]
	LFSR[246] = LFSR[245]
	LFSR[245] = LFSR[244]
	LFSR[244] = LFSR[243]
	LFSR[243] = LFSR[242]
	LFSR[242] = LFSR[241]
	LFSR[241] = LFSR[240]
	LFSR[240] = LFSR[239]
	LFSR[239] = LFSR[238]
	LFSR[238] = LFSR[237]
	LFSR[237] = LFSR[236]
	LFSR[236] = LFSR[235]
	LFSR[235] = LFSR[234]
	LFSR[234] = LFSR[233]
	LFSR[233] = LFSR[232]
	LFSR[232] = LFSR[231]
	LFSR[231] = LFSR[230]
	LFSR[230] = LFSR[229]
	LFSR[229] = LFSR[228]
	LFSR[228] = LFSR[227]
	LFSR[227] = LFSR[226]
	LFSR[226] = LFSR[225]
	LFSR[225] = LFSR[224]
	LFSR[224] = LFSR[223]
	LFSR[223] = LFSR[222]
	LFSR[222] = LFSR[221]
	LFSR[221] = LFSR[220]
	LFSR[220] = LFSR[219]
	LFSR[219] = LFSR[218]
	LFSR[218] = LFSR[217]
	LFSR[217] = LFSR[216]
	LFSR[216] = LFSR[215]
	LFSR[215] = LFSR[214]
	LFSR[214] = LFSR[213]
	LFSR[213] = LFSR[212]
	LFSR[212] = LFSR[211]
	LFSR[211] = LFSR[210]
	LFSR[210] = LFSR[209]
	LFSR[209] = LFSR[208]
	LFSR[208] = LFSR[207]
	LFSR[207] = LFSR[206]
	LFSR[206] = LFSR[205]
	LFSR[205] = LFSR[204]
	LFSR[204] = LFSR[203]
	LFSR[203] = LFSR[202]
	LFSR[202] = LFSR[201]
	LFSR[201] = LFSR[200]
	LFSR[200] = LFSR[199]
	LFSR[199] = LFSR[198]
	LFSR[198] = LFSR[197]
	LFSR[197] = LFSR[196]
	LFSR[196] = LFSR[195]
	LFSR[195] = LFSR[194]
	LFSR[194] = LFSR[193]
	LFSR[193] = LFSR[192]
	LFSR[192] = LFSR[191]
	LFSR[191] = LFSR[190]
	LFSR[190] = LFSR[189]
	LFSR[189] = LFSR[188]
	LFSR[188] = LFSR[187]
	LFSR[187] = LFSR[186]
	LFSR[186] = LFSR[185]
	LFSR[185] = LFSR[184]
	LFSR[184] = LFSR[183]
	LFSR[183] = LFSR[182]
	LFSR[182] = LFSR[181]
	LFSR[181] = LFSR[180]
	LFSR[180] = LFSR[179]
	LFSR[179] = LFSR[178]
	LFSR[178] = LFSR[177]
	LFSR[177] = t2;
	//fmt.Println(LFSR)
}
func main() {
	genKey()
	keySetup()
	for i:=0; i<(1024 *1024*1024)*8; i++ {
		genKeyStream()
	}
	fmt.Println("Done!")
}
