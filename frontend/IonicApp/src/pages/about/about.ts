import { Component } from '@angular/core';

import { NavController } from 'ionic-angular';

@Component({
  selector: 'page-about',
  templateUrl: 'about.html'
})
export class AboutPage {

  constructor(public navCtrl: NavController) {
    //TODO: httpService
    //REST eg. Stats: https://stock.xememah.com/s02/sciana-rest/stats.json
  }

}
