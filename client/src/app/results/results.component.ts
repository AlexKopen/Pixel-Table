import { Component, OnInit } from '@angular/core';
import {DataService} from "../services/data.service";
import {BotState} from "../models/bot-state.model";
import {MarketOrder} from "../models/market-order.model";

@Component({
  selector: 'app-results',
  templateUrl: './results.component.html',
  styleUrls: ['./results.component.scss']
})
export class ResultsComponent implements OnInit {
  botStates: BotState[];
  displayedColumns: string[] = ['Action', 'Price', 'Time'];
  dataSource: MarketOrder[]

  constructor(private dataService: DataService) { }

  ngOnInit(): void {
    this.dataService.botStates$.subscribe((botStates: BotState[]) => {
      this.botStates = botStates;

      if (this.botStates.length > 0) {
        this.dataSource = this.botStates[0].MarketOrders
      }
    })
  }

}
