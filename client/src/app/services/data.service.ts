import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { EngineConfiguration } from '../models/engine-configuration.model';

@Injectable({
  providedIn: 'root'
})
export class DataService {
  constructor(private http: HttpClient) {}

  sendConfig(timestamp: number, symbols: string[]): Observable<any> {
    const config = new EngineConfiguration(symbols, timestamp);
    return this.http.post('/api', config);
  }
}
