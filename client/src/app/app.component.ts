import { Component, OnInit } from '@angular/core';
import { BotState } from './models/bot-state.model';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  botStates: BotState[];

  ngOnInit() {
    const conn = new WebSocket('ws://localhost:8080/ws');

    conn.onclose = evt => {
      console.log('Connection closed');
    };

    conn.onmessage = evt => {
      this.botStates = JSON.parse(evt.data);
    };
  }
}
