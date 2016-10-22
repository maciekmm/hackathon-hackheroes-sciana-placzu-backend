import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Rx';
import 'rxjs/add/operator/map';

/*
	Generated class for the DataService provider.

	See https://angular.io/docs/ts/latest/guide/dependency-injection.html
	for more info on providers and Angular 2 DI.
*/

@Injectable()
export class DataService {
	private rootUrl = 'https://sciana.placzu.pl/';
	
	constructor(public http: Http) {
	}

	topItems: any;
	fetchTop() {
		if(this.topItems)
			return Promise.resolve(this.topItems);
		
		return new Promise(resolve => {
			this.http.get(this.rootUrl+"top?limit=10")
			.map(res => res.json())
			.subscribe(
				data => resolve(data),
				error => console.log(error),
				() => console.log("fetchTop")
			);
		});
	}

	stats: any;
	fetchStats() {
		if(this.stats)
			return Promise.resolve(this.stats);

		return new Promise(resolve => {
			this.http.get(this.rootUrl+"stats")
			.map(res => res.json())
			.subscribe(
				data => resolve(data),
				error => console.log(error),
				() => console.log("fetchStats")
			);
		});
	}

	services: any;
	fetchServices() {
		if(this.services) 
			return Promise.resolve(this.services)

		return new Promise(resolve => {
			this.http.get(this.rootUrl+"services")
			.map(res => res.json())
			.subscribe(
				data => resolve(data),
				error => console.log(error),
				() => console.log("fetchServices")
			);
		});
	}
}
