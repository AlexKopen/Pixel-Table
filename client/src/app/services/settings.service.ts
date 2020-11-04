import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SettingsService {
  selectedSymbols$: BehaviorSubject<string> = new BehaviorSubject('');

  constructor() {}
}
