import { Component } from '@angular/core';

import { TopPage } from '../top/top';
import { AboutPage } from '../about/about';
import { SearchPage } from '../search/search';

@Component({
  templateUrl: 'tabs.html'
})
export class TabsPage {
  // this tells the tabs component which Pages
  // should be each tab's root Page
  tab1Root: any = TopPage;
  tab2Root: any = SearchPage;
  tab3Root: any = AboutPage;

  constructor() {

  }
}
