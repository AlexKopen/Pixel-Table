import { Component, OnInit } from '@angular/core';
import { DataService } from '../services/data.service';
import { BotState } from '../models/bot-state.model';
import { SettingsService } from '../services/settings.service';

@Component({
  selector: 'app-results',
  templateUrl: './results.component.html',
  styleUrls: ['./results.component.scss']
})
export class ResultsComponent implements OnInit {
  botStates: BotState[] = [];
  displayedColumns: string[] = ['Action', 'Price', 'Time'];
  selectedSymbol = '';

  constructor(
    private dataService: DataService,
    private settingsService: SettingsService
  ) {}

  ngOnInit(): void {
    this.dataService.botStates$.subscribe((botStates: BotState[]) => {
      this.botStates = botStates;
    });

    this.settingsService.selectedSymbols$.subscribe((symbol: string) => {
      this.selectedSymbol = symbol;
    });
  }

  viewAllClick(): void {
    this.selectedSymbol = '';
  }

  get selectedBotStates(): BotState[] {
    if (this.selectedSymbol === '') {
      return this.botStates;
    } else {
      return this.botStates.filter((botState: BotState) => {
        return botState.Symbol === this.selectedSymbol;
      });
    }
  }

  get disableViewAll(): boolean {
    return this.selectedSymbol === '';
  }
}
