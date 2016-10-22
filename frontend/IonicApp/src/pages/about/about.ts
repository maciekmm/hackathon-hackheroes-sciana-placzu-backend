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

	constructor(public navCtrl: NavController, public dataService: DataService) {
		dataService.fetchStats().then(data => {
			this.stats = data;
		});
	}
}
