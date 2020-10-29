export class EngineConfiguration {
  Symbols: string[];
  StartingTimeStamp: number;

  constructor(Symbols: string[], StartingTimeStamp: number) {
    this.Symbols = Symbols;
    this.StartingTimeStamp = StartingTimeStamp;
  }
}
