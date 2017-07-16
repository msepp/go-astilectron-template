import { Injectable, NgZone } from '@angular/core';
import { Observable } from 'rxjs/Rx';
import { Subject } from 'rxjs/Subject';
import { ReplaySubject } from 'rxjs/ReplaySubject';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import 'rxjs/add/operator/take';
import 'rxjs/add/operator/filter';
import 'rxjs/add/operator/timeout';

import { Message, MESSAGE_TYPE_ERROR } from './message';

// This is the global window.astilectron object that allows communication between
// render and main process. It needs to be declared like this to avoid reference
// errors.
declare var astilectron: any;

// Astilectron class provides an API for communicating to Astilectron backend
// from TypeScript
@Injectable()
export class Astilectron {
  private _msgId = 0;
  private _isReady: BehaviorSubject<boolean> = new BehaviorSubject(false);
  private _messages: ReplaySubject<Message> = new ReplaySubject(10);

  constructor(private zone: NgZone) {
    // Bind to astilectron-ready event. We'll enable astilectron communication
    // when it becomes available.
    document.addEventListener('astilectron-ready', (e) => {
      this.log('doc.add astilectron-ready');
      this.zone.run(() => this.onReady(e));
    });
  }

  private log(...rest) {
    // console.log('[astilectron]', ...rest);
  }

  // isReady returns an observable that tells if Astilectron is ready for IPC.
  get isReady(): Observable<boolean> {
    return this._isReady.asObservable();
  }

  // onReady is called when global event 'astilectron-ready' is received.
  private onReady(e: Event) {
    this.log('now ready!', e);

    // Setup listener for astilectron messages.
    astilectron.listen((m) => {
      this.zone.run(() => this.onMessage(m));
    });

    // Mark astilectron ready for use.
    this._isReady.next(true);
  }

  // onMessage pushes messages to the message stream as they are received.
  private onMessage(m: Message) {
    this._messages.next(m);
  }

  // send sends an message to backend. Returns an observable that will contain
  // matching result.
  public send(key: string, data: any): Observable<Message> {
    this.log('doing send', key, data);
    // Return error if not ready yet
    if (this._isReady.value === false) {
      const s = new Subject();
      s.error('not ready');
      return s.asObservable();
    }

    const message = {
      key: key,
      data: data,
      id: '' + (++this._msgId)
    };

    astilectron.send(message);

    const sub: Subject<Message> = new Subject();
    this._messages
      .filter((m: Message) => m.id === message.id)
      .timeout(5000)
      .take(1)
      .subscribe(
        (m: Message) => {
          if (m.type === MESSAGE_TYPE_ERROR) {
            sub.error(new Error(m.data));
            return;
          }
          sub.next(m);
        },
        e => sub.error(e),
        () => sub.complete()
      );

    return sub.asObservable();
  }

  // messages returns an observable stream of messages, returning at most 10 latest
  // messages seen before subscription.
  get messages(): Observable<Message> {
    return this._messages.asObservable();
  }
}
