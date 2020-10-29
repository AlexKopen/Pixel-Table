export class EngineConfiguration {
  Symbols: string[];
  EndingTimeStamp: number;

  constructor(Symbols: string[], EndingTimeStamp: number) {
    this.Symbols = Symbols;
    this.EndingTimeStamp = EndingTimeStamp;
  }
}
