// Code generated by oto; DO NOT EDIT.

'use strict';

 
export class CardService {
	
	async card(cardRequest) {
		const headers = {
			'Accept':		'application/json',
			'Accept-Encoding':	'gzip',
			'Content-Type':		'application/json',
		}
		cardRequest = cardRequest || {}
		const response = await fetch('/oto/CardService.Card', {
			method: 'POST',
			headers: headers,
			body: JSON.stringify(cardRequest)
		})
		return response.json().then((json) => {
			if (json.error) {
				throw new Error(json.error)
			}
			return json
		})
	}
	
	async cards(cardsRequest) {
		const headers = {
			'Accept':		'application/json',
			'Accept-Encoding':	'gzip',
			'Content-Type':		'application/json',
		}
		cardsRequest = cardsRequest || {}
		const response = await fetch('/oto/CardService.Cards', {
			method: 'POST',
			headers: headers,
			body: JSON.stringify(cardsRequest)
		})
		return response.json().then((json) => {
			if (json.error) {
				throw new Error(json.error)
			}
			return json
		})
	}
	
}

