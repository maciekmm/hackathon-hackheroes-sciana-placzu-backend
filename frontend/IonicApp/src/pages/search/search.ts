import { Component } from '@angular/core';

import { NavController } from 'ionic-angular';

@Component({
  selector: 'page-search',
  templateUrl: 'search.html'
})
export class SearchPage {

  constructor(public navCtrl: NavController) {
    //TODO: httpService, ActionSheet for filtering
    //REST eg. Providers: https://stock.xememah.com/s02/sciana-rest/provider.json
  }

}
