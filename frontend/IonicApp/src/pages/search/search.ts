import { Component } from '@angular/core';
import { NavController } from 'ionic-angular';
import { DataService } from '../../providers/data-service';
import { ServicePage } from '../service/service';

@Component({
	selector: 'page-search',
	templateUrl: 'search.html',
	providers: [DataService]
})
export class SearchPage {
	public services: Array<any>;
	public copy: Array<any>;

	searchValue: string = "";

	constructor(public navCtrl: NavController, public dataService: DataService) {
		this.dataService.fetchServices().then(data => {
			this.services = data;
			this.copy = data;
		});
	}

	getItems(ev: any) {
		this.services = this.copy;

		let val = ev.target.value;
		if (val && val.trim() != '') {
			this.services = this.services.filter((item) => {
				return (item.name.toLowerCase().indexOf(val.toLowerCase()) > -1);
			})
		}
	}

	goToService(service) {
    	this.navCtrl.push(ServicePage, {service : service});
  	}
}
