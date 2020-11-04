import { MarketOrder } from './market-order.model';
import { OrderCycle } from './order-cycle.model';

export class BotState {
  Symbol: string;
  Active: boolean;
  MarketOrders: MarketOrder[];
  orderCycles: OrderCycle[];
}
