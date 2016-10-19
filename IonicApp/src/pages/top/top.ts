import { Component } from '@angular/core';
import { NavController } from 'ionic-angular';

/*
  Generated class for the Top page.

  See http://ionicframework.com/docs/v2/components/#navigation for more info on
  Ionic pages and navigation.
*/
@Component({
  selector: 'page-top',
  templateUrl: 'top.html'
})
export class Top {

  constructor(public navCtrl: NavController) {}

  ionViewDidLoad() {
    console.log('Hello Top Page');
  }

}
