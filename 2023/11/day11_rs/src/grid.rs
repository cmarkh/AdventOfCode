pub struct Grid<T> {
    pub data: Vec<Vec<T>>,
    pub rows: usize,
    pub columns: usize,
}

impl<T: Clone> Grid<T> {
    pub fn new(rows: usize, columns: usize, default: T) -> Self {
        Grid { data: vec![vec![default; columns]; rows], rows, columns }
    }

    pub fn insert_row(&mut self, index: usize, row: Vec<T>) {
        self.data.insert(index, row);
        self.rows += 1;
    }

    pub fn insert_column(&mut self, index: usize, column: Vec<T>) {
        for (i, row) in self.data.iter_mut().enumerate() {
            row.insert(index, column[i].clone());
        }
        self.columns += 1;
    }

    pub fn get(&self, row: usize, column: usize) -> Option<&T> {
        self.data.get(row)?.get(column)
    }

    pub fn set(&mut self, row: usize, column: usize, value: T) -> anyhow::Result<()> {
        if let Some(r) = self.data.get_mut(row) {
            if let Some(c) = r.get_mut(column) {
                *c = value;
                return Ok(());
            }
        }
        anyhow::bail!("Row or column is out of bounds");
    }
}
