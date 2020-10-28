import { MarketOrderAction } from '../enums/market-order-action.enum';

export class MarketOrder {
  Action: MarketOrderAction;
  Price: number;
  Time: number;
}
