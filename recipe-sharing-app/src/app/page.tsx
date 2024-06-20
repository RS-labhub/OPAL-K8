import pageHome from './pages/index'

// pages/index.tsx

import Head from 'next/head';

export default function Home() {
  return (
    <div className="bg-gray-100 min-h-screen flex flex-col">
      <Head>
        <title>Recipe 🧑‍🍳</title>
        <meta name="description" content="Share and discover recipes from around the world." />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <header className="bg-green-600 text-white p-4 shadow-md">
        <div className="container mx-auto flex justify-between items-center">
          <h1 className="text-2xl font-bold">Recipe 🧑‍🍳</h1>
          <nav>
            <a href="#features" className="px-3 py-2 rounded hover:bg-green-700">Features</a>
            <a href="#about" className="px-3 py-2 rounded hover:bg-green-700">About</a>
            <a href="#contact" className="px-3 py-2 rounded hover:bg-green-700">Contact</a>
          </nav>
        </div>
      </header>

      <main className="flex-grow">
        <section className="bg-white text-center py-16">
          <div className="container mx-auto">
            <h2 className="text-4xl font-bold mb-4">Discover and Share Amazing Recipes</h2>
            <p className="text-gray-600 mb-8">Join our community and explore thousands of recipes from around the world.</p>
            <a href="/add-recipe" className="bg-green-600 text-white px-6 py-3 rounded-md text-lg font-medium hover:bg-green-700">Add Your Recipe</a>
          </div>
        </section>

        <section id="features" className="bg-gray-200 py-16">
          <div className="container mx-auto text-center">
            <h3 className="text-3xl font-bold mb-6">Features</h3>
            <div className="flex flex-wrap justify-center">
              <div className="w-full md:w-1/3 p-4">
                <div className="bg-white p-6 rounded-lg shadow-lg">
                  <h4 className="text-2xl font-bold mb-4">Explore Recipes</h4>
                  <p className="text-gray-600">Browse a wide variety of recipes from different cuisines and dietary preferences.</p>
                </div>
              </div>
              <div className="w-full md:w-1/3 p-4">
                <div className="bg-white p-6 rounded-lg shadow-lg">
                  <h4 className="text-2xl font-bold mb-4">Add Your Own</h4>
                  <p className="text-gray-600">Share your favorite recipes with the community and get feedback.</p>
                </div>
              </div>
              <div className="w-full md:w-1/3 p-4">
                <div className="bg-white p-6 rounded-lg shadow-lg">
                  <h4 className="text-2xl font-bold mb-4">Save Favorites</h4>
                  <p className="text-gray-600">Save your favorite recipes for easy access later.</p>
                </div>
              </div>
            </div>
          </div>
        </section>

        <section id="about" className="bg-white py-16">
          <div className="container mx-auto text-center">
            <h3 className="text-3xl font-bold mb-6">About Us</h3>
            <p className="text-gray-600 max-w-2xl mx-auto">Our mission is to create a platform where food enthusiasts can share and discover new recipes. We believe that cooking brings people together, and our goal is to make it easier for everyone to enjoy delicious, home-cooked meals.</p>
          </div>
        </section>

        <section id="contact" className="bg-gray-200 py-16">
          <div className="container mx-auto text-center">
            <h3 className="text-3xl font-bold mb-6">Contact Us</h3>
            <p className="text-gray-600 max-w-2xl mx-auto">Have any questions or feedback? We'd love to hear from you. Reach out to us at contact@recipeapp.com.</p>
          </div>
        </section>
      </main>

      <footer className="bg-green-600 text-white p-4 text-center">
        <p>&copy; 2024 Recipe App. All rights reserved.</p>
      </footer>
    </div>
  );
}
