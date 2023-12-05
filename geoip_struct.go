/*
** Copyright (C) 2023 - Champ Clark III <dabeave _AT_ gmail.com>
**
** This program is free software; you can redistribute it and/or modify
** it under the terms of the GNU General Public License Version 2 as
** published by the Free Software Foundation.  You may not use, modify or
** distribute this program under any other version of the GNU General
** Public License.
**
** This program is distributed in the hope that it will be useful,
** but WITHOUT ANY WARRANTY; without even the implied warranty of
** MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
** GNU General Public License for more details.
**
** You should have received a copy of the GNU General Public License
** along with this program; if not, write to the Free Software
** Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA 02111-1307, USA.
*/

package main

import ()

/* Taken 2023/11/02 - version 2.1 of the GeoIP API with jsonutil - https://github.com/bashtian/jsonutils */

type GeoIP struct {
	City struct {
		GeonameID int64 `json:"geoname_id"` // 4.569067e+06
		Names     struct {
			En   string `json:"en"`    // Aiken
			Ja   string `json:"ja"`    // エイキン
			Ru   string `json:"ru"`    // Айкен
			ZhCN string `json:"zh-CN"` // 艾肯
		} `json:"names"`
	} `json:"city"`
	Continent struct {
		Code      string `json:"code"`       // NA
		GeonameID int64  `json:"geoname_id"` // 6.255149e+06
		Names     struct {
			De   string `json:"de"`    // Nordamerika
			En   string `json:"en"`    // North America
			Es   string `json:"es"`    // Norteamérica
			Fr   string `json:"fr"`    // Amérique du Nord
			Ja   string `json:"ja"`    // 北アメリカ
			PtBR string `json:"pt-BR"` // América do Norte
			Ru   string `json:"ru"`    // Северная Америка
			ZhCN string `json:"zh-CN"` // 北美洲
		} `json:"names"`
	} `json:"continent"`
	Country struct {
		GeonameID int64  `json:"geoname_id"` // 6.252001e+06
		IsoCode   string `json:"iso_code"`   // US
		Names     struct {
			De   string `json:"de"`    // Vereinigte Staaten
			En   string `json:"en"`    // United States
			Es   string `json:"es"`    // Estados Unidos
			Fr   string `json:"fr"`    // États Unis
			Ja   string `json:"ja"`    // アメリカ
			PtBR string `json:"pt-BR"` // EUA
			Ru   string `json:"ru"`    // США
			ZhCN string `json:"zh-CN"` // 美国
		} `json:"names"`
	} `json:"country"`
	Location struct {
		AccuracyRadius int64   `json:"accuracy_radius"` // 200
		Latitude       float64 `json:"latitude"`        // 33.5077
		Longitude      float64 `json:"longitude"`       // -81.6877
		MetroCode      int64   `json:"metro_code"`      // 520
		TimeZone       string  `json:"time_zone"`       // America/New_York
	} `json:"location"`
	Maxmind struct {
		QueriesRemaining int64 `json:"queries_remaining"` // 16656
	} `json:"maxmind"`
	Postal struct {
		Code string `json:"code"` // 29803
	} `json:"postal"`
	RegisteredCountry struct {
		GeonameID int64  `json:"geoname_id"` // 6.252001e+06
		IsoCode   string `json:"iso_code"`   // US
		Names     struct {
			De   string `json:"de"`    // Vereinigte Staaten
			En   string `json:"en"`    // United States
			Es   string `json:"es"`    // Estados Unidos
			Fr   string `json:"fr"`    // États Unis
			Ja   string `json:"ja"`    // アメリカ
			PtBR string `json:"pt-BR"` // EUA
			Ru   string `json:"ru"`    // США
			ZhCN string `json:"zh-CN"` // 美国
		} `json:"names"`
	} `json:"registered_country"`
	Subdivisions []struct {
		GeonameID int64  `json:"geoname_id"` // 4.59704e+06
		IsoCode   string `json:"iso_code"`   // SC
		Names     struct {
			En   string `json:"en"`    // South Carolina
			Es   string `json:"es"`    // Carolina del Sur
			Fr   string `json:"fr"`    // Caroline du Sud
			Ja   string `json:"ja"`    // サウスカロライナ州
			PtBR string `json:"pt-BR"` // Carolina do Sul
			Ru   string `json:"ru"`    // Южная Каролина
			ZhCN string `json:"zh-CN"` // 南卡罗来纳州
		} `json:"names"`
	} `json:"subdivisions"`
	Traits struct {
		AutonomousSystemNumber       int64  `json:"autonomous_system_number"`       // 3356
		AutonomousSystemOrganization string `json:"autonomous_system_organization"` // LEVEL3
		ConnectionType               string `json:"connection_type"`                // Cable/DSL
		IpAddress                    string `json:"ip_address"`                     // 4.2.2.1
		Isp                          string `json:"isp"`                            // Lumen
		Network                      string `json:"network"`                        // 4.2.2.0/28
		Organization                 string `json:"organization"`                   // Lumen
	} `json:"traits"`
}
