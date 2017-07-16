import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AstilectronReadyResolver } from './astilectron-ready.resolver';
import { Astilectron } from './astilectron';

@NgModule({
  imports: [
    CommonModule
  ],
  declarations: [],
  providers: [
    Astilectron,
    AstilectronReadyResolver
  ]
})
export class AstilectronModule { }
