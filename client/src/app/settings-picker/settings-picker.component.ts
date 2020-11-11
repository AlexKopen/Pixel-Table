import {Component, OnInit} from '@angular/core';
import {DataService} from '../services/data.service';
import {SymbolSelection} from '../models/symbol-selection.model';
import {SYMBOLS} from '../constants/symbols-constant';
import {MatTableDataSource} from '@angular/material/table';
import {SelectionModel} from '@angular/cdk/collections';
import {SettingsService} from '../services/settings.service';
import {BotState} from '../models/bot-state.model';

@Component({
    selector: 'app-settings-picker',
    templateUrl: './settings-picker.component.html',
    styleUrls: ['./settings-picker.component.scss']
})
export class SettingsPickerComponent implements OnInit {
    botStates: BotState[] = [];
    symbolSelections: SymbolSelection[] = [];
    dateStart: Date = new Date();
    dateEnd: Date = new Date();
    displayedColumns: string[] = ['Select', 'Symbol', 'View', 'Profit'];
    dataSource = new MatTableDataSource<SymbolSelection>(this.symbolSelections);
    selection = new SelectionModel<SymbolSelection>(true, []);

    constructor(
        private dataService: DataService,
        private settingsService: SettingsService
    ) {
    }

    ngOnInit(): void {
        SYMBOLS.forEach((symbol: string, index: number) => {
            this.symbolSelections.push(new SymbolSelection(symbol, index + 1));
        });

        this.dataSource.data.forEach(row => this.selection.select(row));

        this.dataService.botStates$.subscribe((botStates: BotState[]) => {
            this.botStates = botStates;
        });
    }

    sendConfig(): void {
        const symbols: string[] = this.selection.selected.map(
            (symbolSelection: SymbolSelection) => {
                return symbolSelection.symbol;
            }
        );

        const currentDate = new Date();
        const endTime =
            currentDate.getFullYear() === this.dateEnd.getFullYear() &&
            currentDate.getMonth() === this.dateEnd.getMonth() &&
            currentDate.getDate() === this.dateEnd.getDate() ? currentDate.getTime() : this.dateEnd.getTime();

        this.dataService
            .sendConfig(symbols, this.dateStart.getTime(), endTime)
            .subscribe(() => {
            });
    }

    /** Whether the number of selected elements matches the total number of rows. */
    isAllSelected(): boolean {
        const numSelected = this.selection.selected.length;
        const numRows = this.dataSource.data.length;
        return numSelected === numRows;
    }

    /** Selects all rows if they are not all selected; otherwise clear selection. */
    masterToggle(): void {
        this.isAllSelected()
            ? this.selection.clear()
            : this.dataSource.data.forEach(row => this.selection.select(row));
    }

    /** The label for the checkbox on the passed row */
    checkboxLabel(row?: SymbolSelection): string {
        if (!row) {
            return `${this.isAllSelected() ? 'select' : 'deselect'} all`;
        }
        return `${this.selection.isSelected(row) ? 'deselect' : 'select'} row ${
            row.position + 1
        }`;
    }

    viewClick(symbol: string): void {
        this.settingsService.selectedSymbols$.next(symbol);
    }

    totalProfit(symbol: string): number {
        const botState = this.botStates.find((currentBotState: BotState) => {
            return currentBotState.Symbol === symbol;
        });

        return botState !== undefined
            ? this.dataService.totalProfit(botState.orderCycles)
            : 0;
    }

    get totalProfitAll(): number {
        let totalProfit = 0;

        this.symbolSelections.forEach((symbolSelection: SymbolSelection) => {
            totalProfit += this.totalProfit(symbolSelection.symbol);
        });

        return totalProfit;
    }
}
