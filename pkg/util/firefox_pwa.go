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

type FIREFOX_PWA struct {
	Profiles  map[string]Profile `json:"profiles"`
	Sites     map[string]Site    `json:"sites"`
	Arguments []interface{}      `json:"arguments"`
	Variables struct{}           `json:"variables"`
	Config    struct {
		AlwaysPatch          bool `json:"always_patch"`
		RuntimeEnableWayland bool `json:"runtime_enable_wayland"`
		RuntimeUseXinput2    bool `json:"runtime_use_xinput2"`
		RuntimeUsePortals    bool `json:"runtime_use_portals"`
	} `json:"config"`
}

type Profile struct {
	Ulid        string   `json:"ulid"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Sites       []string `json:"sites"`
}

type Site struct {
	Ulid        string        `json:"ulid"`
	Profile     string        `json:"profile"`
	Config      SiteConfig    `json:"config"`
	Icons       []Icon        `json:"icons"`
	Screenshots []interface{} `json:"screenshots"`
}

type SiteConfig struct {
	Name                    string        `json:"name"`
	Description             string        `json:"description"`
	StartURL                string        `json:"start_url"`
	DocumentURL             string        `json:"document_url"`
	ManifestURL             string        `json:"manifest_url"`
	Categories              interface{}   `json:"categories"`
	Keywords                interface{}   `json:"keywords"`
	EnabledURLHandlers      []interface{} `json:"enabled_url_handlers"`
	EnabledProtocolHandlers []interface{} `json:"enabled_protocol_handlers"`
	CustomProtocolHandlers  []interface{} `json:"custom_protocol_handlers"`
}

type SiteManifest struct {
	StartURL                  string         `json:"start_url"`
	Scope                     string         `json:"scope"`
	Name                      string         `json:"name"`
	ShortName                 string         `json:"short_name"`
	Categories                []interface{}  `json:"categories"`
	Keywords                  []interface{}  `json:"keywords"`
	Dir                       string         `json:"dir"`
	Display                   string         `json:"display"`
	Orientation               string         `json:"orientation"`
	BackgroundColor           string         `json:"background_color"`
	ThemeColor                string         `json:"theme_color"`
	PreferRelatedApplications bool           `json:"prefer_related_applications"`
	RelatedApplications       []interface{}  `json:"related_applications"`
	ProtocolHandlers          []interface{}  `json:"protocol_handlers"`
	Shortcuts                 []SiteShortcut `json:"shortcuts"`
	Icons                     []Icon         `json:"icons"`
	Screenshots               []interface{}  `json:"screenshots"`
}

type SiteShortcut struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Icons []Icon `json:"icons"`
}

type Icon struct {
	Src     string `json:"src"`
	Type    string `json:"type"`
	Sizes   string `json:"sizes"`
	Purpose string `json:"purpose"`
}
