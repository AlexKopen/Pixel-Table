import { MarketOrder } from './market-order.model';

export class BotState {
  Symbol: string;
  Active: boolean;
  MarketOrders: MarketOrder[];
}
