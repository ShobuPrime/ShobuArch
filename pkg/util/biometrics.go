/*
* Automated Arch Linux Tools
* Copyright (C) 2022  Anthony Dardano (ShobuPrime)
*
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* ShobuArch is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package util

type BioSupportedDevices struct {
	Face        []string
	Fingerprint []string
	SecurityKey []string
}

func BiometricIDs() *BioSupportedDevices {

	bsd := &BioSupportedDevices{}

	bsd.Face = []string{
		`04f2:b5ac`,
		`04f2:b5cf`,
		`04f2:b6d0`,
	}

	// List of libfprint supported devices: https://fprint.freedesktop.org/supported-devices.html
	bsd.Fingerprint = []string{
		`045e:00bb`,
		`045e:00bc`,
		`045e:00bd`,
		`045e:00ca`,
		`0483:2015`,
		`0483:2016`,
		`04f3:0903`,
		`04f3:0907`,
		`04f3:0c01`,
		`04f3:0c02`,
		`04f3:0c03`,
		`04f3:0c04`,
		`04f3:0c05`,
		`04f3:0c06`,
		`04f3:0c07`,
		`04f3:0c08`,
		`04f3:0c09`,
		`04f3:0c0a`,
		`04f3:0c0b`,
		`04f3:0c0c`,
		`04f3:0c0d`,
		`04f3:0c0e`,
		`04f3:0c0f`,
		`04f3:0c10`,
		`04f3:0c11`,
		`04f3:0c12`,
		`04f3:0c13`,
		`04f3:0c14`,
		`04f3:0c15`,
		`04f3:0c16`,
		`04f3:0c17`,
		`04f3:0c18`,
		`04f3:0c19`,
		`04f3:0c1a`,
		`04f3:0c1b`,
		`04f3:0c1c`,
		`04f3:0c1d`,
		`04f3:0c1e`,
		`04f3:0c1f`,
		`04f3:0c20`,
		`04f3:0c21`,
		`04f3:0c22`,
		`04f3:0c23`,
		`04f3:0c24`,
		`04f3:0c25`,
		`04f3:0c26`,
		`04f3:0c27`,
		`04f3:0c28`,
		`04f3:0c29`,
		`04f3:0c2a`,
		`04f3:0c2b`,
		`04f3:0c2c`,
		`04f3:0c2d`,
		`04f3:0c2e`,
		`04f3:0c2f`,
		`04f3:0c30`,
		`04f3:0c31`,
		`04f3:0c32`,
		`04f3:0c33`,
		`04f3:0c3d`,
		`04f3:0c42`,
		`04f3:0c4b`,
		`04f3:0c4d`,
		`04f3:0c4f`,
		`04f3:0c58`,
		`04f3:0c63`,
		`04f3:0c6e`,
		`04f3:0c7d`,
		`04f3:0c7e`,
		`04f3:0c82`,
		`04f3:0c88`,
		`04f3:0c8c`,
		`04f3:0c8d`,
		`05ba:0007`,
		`05ba:0008`,
		`05ba:000a`,
		`061a:0110`,
		`06cb:00bd`,
		`06cb:00c2`,
		`06cb:00df`,
		`06cb:00f0`,
		`06cb:00f9`,
		`06cb:00fc`,
		`06cb:0100`,
		`06cb:0103`,
		`06cb:0104`,
		`06cb:0123`,
		`06cb:0126`,
		`06cb:0129`,
		`06cb:015f`,
		`06cb:0168`,
		`08ff:1600`,
		`08ff:1660`,
		`08ff:1680`,
		`08ff:1681`,
		`08ff:1682`,
		`08ff:1683`,
		`08ff:1684`,
		`08ff:1685`,
		`08ff:1686`,
		`08ff:1687`,
		`08ff:1688`,
		`08ff:1689`,
		`08ff:168a`,
		`08ff:168b`,
		`08ff:168c`,
		`08ff:168d`,
		`08ff:168e`,
		`08ff:168f`,
		`08ff:2500`,
		`08ff:2550`,
		`08ff:2580`,
		`08ff:2660`,
		`08ff:2680`,
		`08ff:2681`,
		`08ff:2682`,
		`08ff:2683`,
		`08ff:2684`,
		`08ff:2685`,
		`08ff:2686`,
		`08ff:2687`,
		`08ff:2688`,
		`08ff:2689`,
		`08ff:268a`,
		`08ff:268b`,
		`08ff:268c`,
		`08ff:268d`,
		`08ff:268e`,
		`08ff:268f`,
		`08ff:2691`,
		`08ff:2810`,
		`08ff:5731`,
		`138a:0001`,
		`138a:0005`,
		`138a:0008`,
		`138a:0010`,
		`138a:0011`,
		`138a:0015`,
		`138a:0017`,
		`138a:0018`,
		`138a:0050`,
		`138a:0091`,
		`147e:1000`,
		`147e:1001`,
		`147e:2016`,
		`147e:2020`,
		`147e:3001`,
		`1c7a:0570`,
		`1c7a:0571`,
		`1c7a:0603`,
		`27c6:5840`,
		`27c6:6094`,
		`27c6:609c`,
		`27c6:60a2`,
		`27c6:631c`,
		`27c6:634c`,
		`27c6:6384`,
		`27c6:639c`,
		`27c6:63ac`,
		`27c6:63bc`,
		`27c6:63cc`,
		`27c6:6496`,
		`27c6:6584`,
		`27c6:658c`,
		`27c6:6592`,
		`27c6:6594`,
		`27c6:659a`,
		`27c6:659c`,
		`27c6:6a94`,
		`298d:1010`,
		`5501:08ff`,
	}

	// List of FIDO devs: https://github.com/Yubico/libfido2/blob/main/udev/fidodevs
	fido_vendors := map[string]string{
		"STMicro":     `0483`, // STMicroelectronics
		"Infineon":    `058b`, // Infineon Technologies
		"Synaptics":   `06cb`, // Synaptics Inc.
		"Feitian":     `096e`, // Feitian Technologies Co., Ltd.
		"Yubico":      `1050`, // Yubico AB
		"Silicon":     `10c4`, // Silicon Laboratories, Inc.
		"pid.codes":   `1209`, // pid.codes
		"Google":      `18d1`, // Google Inc.
		"VASCO":       `1a44`, // VASCO Data Security NV
		"OpenMoko":    `1d50`, // OpenMoko, Inc.
		"NEOWAVE":     `1e0d`, // NEOWAVE
		"Excelsecu":   `1ea8`, // Shenzhen Excelsecu Data Technology Co., Ltd
		"NXP":         `1fc9`, // NXP Semiconductors
		"ClayLogic":   `20a0`, // Clay Logic
		"Aladdin":     `24dc`, // Aladdin Software Security R.D.
		"Plug‐up":     `2581`, // Plug‐up
		"Bluink":      `2abe`, // Bluink Ltd
		"LEDGER":      `2c97`, // LEDGER
		"Hypersecu":   `2ccf`, // Hypersecu Information Systems, Inc.
		"eWBM":        `311f`, // eWBM Co., Ltd. (TrustKey)
		"GoTrustID":   `32a3`, // GoTrustID Inc.
		"Unknown":     `4c4d`, // Unknown vendor
		"SatoshiLabs": `534c`, // SatoshiLabs
	}

	for vendor, id := range fido_vendors {

		switch vendor {
		case "STMicro":
			products := []string{
				`a2ac`, // ellipticSecure MIRKey
				`a2ca`, // Unknown product
				`cdab`, // Unknown product
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "Infineon":
			products := []string{
				`022d`, // Infineon FIDO
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "Synaptics":
			products := []string{
				`0088`, // Kensington Verimark
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "Feitian":
			products := []string{
				`0850`, // FS ePass FIDO
				`0852`, // Unknown product
				`0853`, // Unknown product
				`0854`, // Unknown product
				`0856`, // Unknown product
				`0858`, // Unknown product
				`085a`, // FS MultiPass FIDO U2F
				`085b`, // Unknown product
				`085d`, // Unknown product
				`0866`, // BioPass FIDO2 K33
				`0867`, // BioPass FIDO2 K43
				`0880`, // Hypersecu HyperFIDO
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "Yubico":
			products := []string{
				`0113`, // YubiKey NEO FIDO
				`0114`, // YubiKey NEO OTP+FIDO
				`0115`, // YubiKey NEO FIDO+CCID
				`0116`, // YubiKey NEO OTP+FIDO+CCID
				`0120`, // Security Key by Yubico
				`0121`, // Unknown product
				`0200`, // Gnubby U2F
				`0402`, // YubiKey 4 FIDO
				`0403`, // YubiKey 4 OTP+FIDO
				`0406`, // YubiKey 4 FIDO+CCID
				`0407`, // YubiKey 4 OTP+FIDO+CCID
				`0410`, // YubiKey Plus
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "Silicon":
			products := []string{
				`8acf`, // U2F Zero
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "pid.codes":
			products := []string{
				`5070`, // SoloKeys SoloHacker
				`50b0`, // SoloKeys SoloBoot
				`53c1`, // SatoshiLabs TREZOR
				`beee`, // SoloKeys v2
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "Google":
			products := []string{
				`5026`, // Google Titan U2F
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "VASCO":
			products := []string{
				`00bb`, // VASCO SecureClick
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "OpenMoko":
			products := []string{
				`60fc`, // OnlyKey (FIDO2/U2F)
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "NEOWAVE":
			products := []string{
				`f1ae`, // Neowave Keydo AES
				`f1d0`, // Neowave Keydo

			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "Excelsecu":
			products := []string{
				`f025`, // Thethis Key
				`fc25`, // ExcelSecu FIDO2 Security Key
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "NXP":
			products := []string{
				`f143`, // GoTrust Idem Key
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "ClayLogic":
			products := []string{
				`4287`, // Nitrokey FIDO U2F
				`42b1`, // Nitrokey FIDO2
				`42b2`, // Nitrokey 3C NFC
				`42b3`, // Safetech SafeKey
				`42d4`, // CanoKey
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "Alladdin":
			products := []string{
				`0101`, // JaCarta U2F
				`0501`, // JaCarta U2F
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "Plug-up":
			products := []string{
				`f1d0`, // Happlink Security Key
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "Bluink":
			products := []string{
				`1002`, // Bluink Key
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "LEDGER":
			products := []string{
				`0000`, // Ledger Blue
				`0001`, // Ledger Nano S Old firmware
				`0004`, // Ledger Nano X Old firmware
				`0011`, // Ledger Blue
				`0015`, // Ledger Blue Legacy
				`1011`, // Ledger Nano S
				`1015`, // Ledger Nano S Legacy
				`4011`, // Ledger Nano X
				`4015`, // Ledger Nano X Legacy
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "HyperSecu":
			products := []string{
				`0880`, // Hypersecu HyperFIDO
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "eWBM":
			products := []string{
				`4a2a`, // TrustKey Solutions FIDO2 G310H/G320H
				`4a1a`, // TrustKey Solutions FIDO2 G310
				`4c2a`, // TrustKey Solutions FIDO2 G320
				`5c2f`, // eWBM FIDO2 Goldengate G500
				`a6e9`, // TrustKey Solutions FIDO2 T120
				`a7f9`, // TrustKey Solutions FIDO2 T110
				`f47c`, // eWBM FIDO2 Goldengate G450
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "GoTrust":
			products := []string{
				`3201`, // Idem Key
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "Unknown":
			products := []string{
				`f703`, // Longmai mFIDO
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		case "SatoshiLabs":
			products := []string{
				`0001`, // SatoshiLabs TREZOR
			}
			CombineIDs(&id, &products, &bsd.SecurityKey)
		}
	}

	return bsd
}
