import { Component } from '@angular/core';
import { NavController } from 'ionic-angular';
import { DataService } from '../../providers/data-service' 

@Component({
	selector: 'page-home',
	templateUrl: 'home.html',
  	providers: [DataService]
})
export class HomePage {
	public topItems: Array<any>;

	constructor(public navCtrl: NavController, public dataService: DataService) {
		dataService.fetchTop().then(data => {
			this.topItems = data;
		});
	}
}
