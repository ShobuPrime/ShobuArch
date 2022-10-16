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
}

func BiometricIDs() *BioSupportedDevices {

	bsd := &BioSupportedDevices{}

	bsd.Face = []string{
		`04f2:b5ac`,
		`04f2:b5cf`,
		`04f2:b6d0`,
	}

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

	return bsd
}
