package net

var maskBitMap = map[int]string{
	1:  "128.0.0.0",
	2:  "192.0.0.0",
	3:  "224.0.0.0",
	4:  "240.0.0.0",
	5:  "248.0.0.0",
	6:  "252.0.0.0",
	7:  "254.0.0.0",
	8:  "255.0.0.0",
	9:  "255.128.0.0",
	10: "255.192.0.0",
	11: "255.224.0.0",
	12: "255.240.0.0",
	13: "255.248.0.0",
	14: "255.252.0.0",
	15: "255.254.0.0",
	16: "255.255.0.0",
	17: "255.255.128.0",
	18: "255.255.192.0",
	19: "255.255.224.0",
	20: "255.255.240.0",
	21: "255.255.248.0",
	22: "255.255.252.0",
	23: "255.255.254.0",
	24: "255.255.255.0",
	25: "255.255.255.128",
	26: "255.255.255.192",
	27: "255.255.255.224",
	28: "255.255.255.240",
	29: "255.255.255.248",
	30: "255.255.255.252",
	31: "255.255.255.254",
	32: "255.255.255.255",
}

// GetMask 根据掩码位获取掩码
func GetMask(maskBit int) string {
	return maskBitMap[maskBit]
}

// GetMaskBit 根据掩码获取掩码位
func GetMaskBit(mask string) int {
	for maskBit, value := range maskBitMap {
		if value == mask {
			return maskBit
		}
	}
	return 0 // 返回0表示不合法
}

// IsMaskValid 判断掩码是否合法
func IsMaskValid(mask string) bool {
	return GetMaskBit(mask) != 0
}

// IsMaskBitValid 判断掩码位是否合法
func IsMaskBitValid(maskBit int) bool {
	return GetMask(maskBit) != ""
}
