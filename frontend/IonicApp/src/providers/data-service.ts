import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import 'rxjs/add/operator/map';

/*
	Generated class for the DataService provider.

	See https://angular.io/docs/ts/latest/guide/dependency-injection.html
	for more info on providers and Angular 2 DI.
*/

export class TopItem {

}

@Injectable()
export class DataService {
	private rootUrl = 'https://stock.xememah.com/s02/sciana-rest/';
	items: any;

	constructor(public http: Http) {
		console.log('init');
	}

	fetchTop() {
		this.http.get(this.rootUrl+"top.json").subscribe(data => {
			this.items = data.json();
			console.log(this.items);
		});
	}
}
