export class EngineConfiguration {
  Symbols: string[];
  StartingTimeStamp: number;
  EndingTimeStamp: number;

  constructor(Symbols: string[], StartingTimeStamp: number, EndingTimeStamp: number) {
    this.Symbols = Symbols;
    this.StartingTimeStamp = StartingTimeStamp;
    this.EndingTimeStamp = EndingTimeStamp;
  }
}
