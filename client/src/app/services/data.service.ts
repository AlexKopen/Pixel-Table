import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject, Observable } from 'rxjs';
import { EngineConfiguration } from '../models/engine-configuration.model';
import { BotState } from '../models/bot-state.model';

@Injectable({
  providedIn: 'root'
})
export class DataService {
  botStates$: BehaviorSubject<BotState[]> = new BehaviorSubject<BotState[]>([]);
  constructor(private http: HttpClient) {}

  sendConfig(timestamp: number, symbols: string[]): Observable<any> {
    const config = new EngineConfiguration(symbols, timestamp);
    return this.http.post('/api', config);
  }

  connectToSocket(): void {
    const conn = new WebSocket('ws://localhost:8080/ws');

    conn.onclose = evt => {
      console.log('Connection closed');
    };

    conn.onmessage = evt => {
      this.botStates$.next(JSON.parse(evt.data));
    };
  }
}
