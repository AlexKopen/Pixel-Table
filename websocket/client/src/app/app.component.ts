import {Component, OnInit} from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit{
  title = 'pixel-table-client';

  ngOnInit() {
    const conn = new WebSocket("ws://localhost:8080/ws?lastMod=16408325606d65e9");
    conn.onclose = function(evt) {
      console.log('Connection closed')
    }
    conn.onmessage = function(evt) {
      console.log('file updated');
      console.log(evt.data)
    }
  }
}
