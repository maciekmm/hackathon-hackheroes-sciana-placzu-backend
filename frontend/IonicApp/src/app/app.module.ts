import { NgModule, Pipe } from '@angular/core';
import { IonicApp, IonicModule } from 'ionic-angular';
import { MyApp } from './app.component';

import { AboutPage } from '../pages/about/about';
import { HomePage } from '../pages/home/home';
import { TabsPage } from '../pages/tabs/tabs';
import { SearchPage } from '../pages/search/search';
import { ServicePage } from '../pages/service/service';
//import { PhonePipe } from '../pages/service/phone-pipe';
import { DataService } from '../providers/data-service';

@NgModule({
  declarations: [
    MyApp,
    AboutPage,
    HomePage,
    TabsPage,
    SearchPage,
    ServicePage //,
                //PhonePipe
  ],
  imports: [
    IonicModule.forRoot(MyApp)
  ],
  bootstrap: [IonicApp],
  entryComponents: [
    MyApp,
    AboutPage,
    HomePage,
    TabsPage,
    SearchPage,
    ServicePage //,
                //PhonePipe
  ],
  providers: [DataService]
})
export class AppModule {}
