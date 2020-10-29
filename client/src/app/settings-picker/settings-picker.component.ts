import { Component, OnInit } from '@angular/core';
import { DataService } from '../services/data.service';
import { SymbolSelection } from '../models/symbol-selection.model';
import { SYMBOLS } from '../constants/symbols-constant';
import { MatTableDataSource } from '@angular/material/table';
import { SelectionModel } from '@angular/cdk/collections';

@Component({
  selector: 'app-settings-picker',
  templateUrl: './settings-picker.component.html',
  styleUrls: ['./settings-picker.component.scss']
})
export class SettingsPickerComponent implements OnInit {
  symbolSelections: SymbolSelection[] = [];
  dateSelection: Date = new Date();
  displayedColumns: string[] = ['select', 'symbol'];
  dataSource = new MatTableDataSource<SymbolSelection>(this.symbolSelections);
  selection = new SelectionModel<SymbolSelection>(true, []);

  constructor(private dataService: DataService) {}

  ngOnInit(): void {
    SYMBOLS.forEach((symbol: string, index: number) => {
      this.symbolSelections.push(new SymbolSelection(symbol, index + 1));
    });

    this.dataSource.data.forEach(row => this.selection.select(row));
  }

  sendConfig(): void {
    const symbols: string[] = this.selection.selected.map(
      (symbolSelection: SymbolSelection) => {
        return symbolSelection.symbol;
      }
    );

    this.dataService
      .sendConfig(this.dateSelection.getTime(), symbols)
      .subscribe(() => {
        console.log('config sent');
      });
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.dataSource.data.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
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
}
