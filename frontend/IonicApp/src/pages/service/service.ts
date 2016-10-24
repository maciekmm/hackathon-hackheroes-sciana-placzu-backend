import { Component } from '@angular/core';
import { NavController, NavParams } from 'ionic-angular';
import { DataService } from '../../providers/data-service';

@Component({
	selector: 'page-service',
	templateUrl: 'service.html',
	providers: [DataService]
})
export class ServicePage {
	public service: any;
	public providers: Array<any>;
	public dataService: DataService;

	voivodeship: string = localStorage["voivodeship-filter"] || "";

	constructor(public navCtrl: NavController, public params: NavParams, public dataServ: DataService) {
		this.service = params.get('service');
		this.dataService = dataServ;
		this.getItems();
	}

	getItems() {
		this.dataService.fetchSearch({name: this.service.name, voivodeship: this.voivodeship}).then(data => {
			this.providers = data;
		});
	}

	updateFilter() {
		localStorage["voivodeship-filter"] = this.voivodeship;
		this.getItems();
	}
}