import { Injectable } from '@angular/core';
import { Resolve, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import { Observable, Subject } from 'rxjs/Rx';

import { Astilectron } from './astilectron';

// AstilectronReadyResolver waits until Astilectron ready state changes to true.
@Injectable()
export class AstilectronReadyResolver implements Resolve<any> {
  constructor(private asti: Astilectron) {}
  resolve(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<any> | Promise<any> | any {
    return this.asti.isReady.filter((s) => s === true).take(1);
  }
}
