import { Component } from '@angular/core';
import { NavController } from 'ionic-angular';
import { DataService } from './providers/data.service' 

@Component({
  selector: 'page-home',
  templateUrl: 'home.html'
})
export class HomePage {

  constructor(public navCtrl: NavController) {
    //TODO: httpService
    //REST eg. Top: https://stock.xememah.com/s02/sciana-rest/top.json
  }

}
