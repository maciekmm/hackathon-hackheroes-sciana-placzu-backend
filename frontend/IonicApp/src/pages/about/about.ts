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
		this.randomizeAuthors();
	}

	ionViewWillEnter() {
		this.randomizeAuthors();
	}

	shuffle(array) {
		var i = 0, j = 0, temp = null;

		for (i = array.length - 1; i > 0; i -= 1) {
			j = Math.floor(Math.random() * (i + 1));
			temp = array[i];
			array[i] = array[j];
			array[j] = temp;
		}
	}

	randomizeAuthors() {
		this.authors = [
			{mail: "marek@kochanow.ski", desc: "Front-end, aplikacja mobilna"},
			{mail: "business@maciekmm.net", desc: "Back-end, API, parsowanie danych"}
		]
		this.shuffle(this.authors);
	}
}
