import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject, Observable } from 'rxjs';
import { EngineConfiguration } from '../models/engine-configuration.model';
import { BotState } from '../models/bot-state.model';
import { MarketOrder } from '../models/market-order.model';
import { OrderCycle } from '../models/order-cycle.model';

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
      this.fillOrderCycles();
    };
  }

  fillOrderCycles(): void {
    this.botStates$.value.forEach((botState: BotState) => {
      botState.orderCycles = this.generateOrderCycles(botState.MarketOrders);
    });
  }

  private generateOrderCycles(marketOrders: MarketOrder[]): OrderCycle[] {
    const orderCycles: OrderCycle[] = [];
    const sortedMarketOrders = marketOrders.sort(this.sortMarketOrdersCompare);

    for (let i = 0; i < sortedMarketOrders.length; i += 2) {
      if (i + 1 <= sortedMarketOrders.length - 1) {
        const orderCycle = new OrderCycle();

        orderCycle.startingPrice = sortedMarketOrders[i].Price;
        orderCycle.startingTime = sortedMarketOrders[i].Time;
        orderCycle.endingPrice = sortedMarketOrders[i + 1].Price;
        orderCycle.endingTime = sortedMarketOrders[i + 1].Time;
        orderCycle.profit =
          (orderCycle.endingPrice - orderCycle.startingPrice) /
          orderCycle.startingPrice;

        orderCycles.push(orderCycle);
      }
    }

    return orderCycles;
  }

  private sortMarketOrdersCompare(a: MarketOrder, b: MarketOrder): number {
    return a.Time > b.Time ? 1 : -1;
  }

  totalProfit(orderCycles: OrderCycle[]): number {
    return orderCycles.reduce((previousValue, currentOrder) => {
      return previousValue + currentOrder.profit;
    }, 0);
  }
}
