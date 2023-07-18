// Step 1: Create or open a database
const request = indexedDB.open('myDatabase', 1);

// Step 2: Handle database upgrades
request.onupgradeneeded = function(event) {
  const db = event.target.result;
  const objectStore = db.createObjectStore('myObjectStore', { keyPath: 'id' });
  objectStore.createIndex('nameIndex', 'name', { unique: true });
};

// Step 3: Handle database connections
request.onsuccess = function(event) {
  const db = event.target.result;

  // Step 4: Perform database operations
  const transaction = db.transaction('myObjectStore', 'readwrite');
  const objectStore = transaction.objectStore('myObjectStore');

  // Add data
//   objectStore.add({ id: 3, name: 'Somesh Prajapati' });
  
  // Retrieve data
  const getRequest = objectStore.get(2);
  getRequest.onsuccess = function(event) {
    const data = event.target.result;
    console.log(data);
  };

  // Step 5: Handle transactions
  transaction.oncomplete = function(event) {
    console.log('Transaction completed.');
  };

  transaction.onerror = function(event) {
    console.error('Transaction error:', event.target.error);
  };
};

request.onerror = function(event) {
  console.error('Database error:', event.target.error);
};
