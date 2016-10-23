import { Component } from '@angular/core';
import { NavController, NavParams } from 'ionic-angular';
import { DataService } from '../../providers/data-service';
import { PhonePipe } from 'phone-pipe';

@Component({
	selector: 'page-service',
	templateUrl: 'service.html',
	providers: [DataService]
})
export class ServicePage {
	public service: any;
	public providers: Array<any>;

	constructor(public navCtrl: NavController, public params: NavParams, public dataService: DataService) {
		this.service = params.get('service');

		console.log(this.service.name);
		
		dataService.fetchSearch(this.service.name).then(data => {
			this.providers = data;
		});
	}

	ionViewDidLoad() {
		console.log('Hello Service Page');
	}

}