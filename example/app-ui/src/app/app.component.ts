import { Component, OnInit } from '@angular/core';
import { Astilectron, Message, REQUEST_APP_VERSIONS } from './astilectron';


import { AppVersion } from './app-version';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.less']
})
export class AppComponent implements OnInit {
  public title = 'go-astilectron template project';
  public versions: AppVersion = null;

  constructor(
    private asti: Astilectron
  ) {}

  public ngOnInit() {
    this.asti.isReady.filter(v => v === true).take(1).subscribe(() => this.getVersions());
  }

  private getVersions() {
    this.asti.send(REQUEST_APP_VERSIONS, null).subscribe(
      (m: Message) => {
        const v: AppVersion = Object.assign(new AppVersion, m.data);
        this.versions = v;
      },
      (e: Error) => {
        console.log('whoops:', e.message);
      }
    );
  }
}
