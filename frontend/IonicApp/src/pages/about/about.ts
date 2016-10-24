import { Component } from '@angular/core';
import { NavController } from 'ionic-angular';
import { DataService } from '../../providers/data-service';

@Component({
	selector: 'page-about',
	templateUrl: 'about.html',
	providers: [DataService]
})
export class AboutPage {
	public stats: Array<any>;
	public authors: Array<any>;

	constructor(public navCtrl: NavController, public dataService: DataService) {
		dataService.fetchStats().then(data => {
			this.stats = data;
		});
		this.authors = [
			{mail: "marek@kochanow.ski", desc: "Front-end, aplikacja mobilna"},
			{mail: "hackheroes@maciekmm.net", desc: "Back-end, API, parsowanie danych"}
		]
	}
}