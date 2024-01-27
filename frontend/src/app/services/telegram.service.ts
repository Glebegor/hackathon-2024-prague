import {Inject, Injectable} from '@angular/core';
import {DOCUMENT} from "@angular/common";

interface TgButton {
  show(): void;
  hide(): void;
  setText(text:String): void;
  onClick(fn:Function):void;
  offClick(fn:Function):void;
}

@Injectable({
  providedIn: 'root'
})
  
export class TelegramService {

  private window;
  tg;
  constructor(
    @Inject(DOCUMENT) private _document: Document) {
    this.window = this._document.defaultView;
    this.tg = this.window.Telegram.WebApp;
  }

  get MainButton(): TgButton {
   return this.tg?.MainButton;
   }

  get BackButton(): TgButton{
    return this.tg?.BackButton;
  }

  ready(){
    this.tg.ready();
  }
  
  
}

