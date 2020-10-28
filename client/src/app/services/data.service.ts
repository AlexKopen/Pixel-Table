import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {EngineConfiguration} from "../models/engine-configuration.model";

@Injectable({
  providedIn: 'root'
})
export class DataService {

  constructor(private http: HttpClient) { }

  sendConfig(): Observable<any> {
    const config = new EngineConfiguration(
      ['LTC', 'XMR', 'BTC'],
      new Date().getTime()
    )
    return this.http.post('/api', config)
  }
}
